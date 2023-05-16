package config

import (
	"bytes"
	"context"
	"testing"

	"github.com/Interhyp/metadata-service/acorns/config"
	"github.com/Interhyp/metadata-service/docs"
	auconfigenv "github.com/StephanHCB/go-autumn-config-env"
	goauzerolog "github.com/StephanHCB/go-autumn-logging-zerolog"
	librepo "github.com/StephanHCB/go-backend-service-common/acorns/repository"
	libconfig "github.com/StephanHCB/go-backend-service-common/repository/config"
	"github.com/StephanHCB/go-backend-service-common/repository/logging"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
)

const basedir = "../../../test/resources/"

func tstYamlRead(t *testing.T, filename string, expectedMsgPart string) {
	cut := New().(*libconfig.ConfigImpl)
	auconfigenv.LocalConfigFileName = basedir + filename
	err := cut.Read()
	if expectedMsgPart == "" {
		require.Nil(t, err)
	} else {
		require.NotNil(t, err)
		require.Contains(t, err.Error(), expectedMsgPart)
	}
}

func TestYamlRead_MissingFile(t *testing.T) {
	docs.Description("the local configuration file is optional")
	tstYamlRead(t, "not-there.yaml", "")
}

func TestYamlRead_InvalidSyntax(t *testing.T) {
	docs.Description("the local configuration must be correct yaml syntax")
	tstYamlRead(t, "invalid-config-syntax.yaml", "error parsing local configuration flat yaml file")
}

func tstSetupCutAndLogRecorder(t *testing.T, configfile string) (librepo.Configuration, error) {
	cut := New().(librepo.Configuration)

	// --- simulate auacornapi.Acorn Assemble phase for just the configuration

	auconfigenv.LocalConfigFileName = basedir + configfile
	err := cut.Read()
	require.Nil(t, err)

	// --- simulate auacornapi.Acorn Setup phase for just the configuration, adding a mock log recorder

	// set up log recorder
	logRecorder := logging.New().(librepo.Logging)
	goauzerolog.RecordedLogForTesting = new(bytes.Buffer)
	logRecorder.(*logging.LoggingImpl).SetupForTesting()

	cut.(*libconfig.ConfigImpl).Logging = logRecorder

	ctx := log.Logger.WithContext(context.Background())
	err = cut.Validate(ctx)

	cut.(*libconfig.ConfigImpl).ObtainPredefinedValues()
	cut.(*libconfig.ConfigImpl).CustomConfiguration.Obtain(auconfigenv.Get)

	return cut, err
}

func TestValidate_LotsOfErrors(t *testing.T) {
	docs.Description("validation of configuration values works")

	_, err := tstSetupCutAndLogRecorder(t, "invalid-config-values.yaml")

	require.NotNil(t, err)
	require.Contains(t, err.Error(), "some configuration values failed to validate or parse. There were 27 error(s). See details above")

	actualLog := goauzerolog.RecordedLogForTesting.String()

	expectedPart := "\"message\":\"failed to validate configuration field ENVIRONMENT: must match ^(feat|"
	require.Contains(t, actualLog, expectedPart)

	expectedPart2 := "\"message\":\"failed to validate configuration field SERVER_PORT: value 122834 is out of range [1024..65535]"
	require.Contains(t, actualLog, expectedPart2)

	expectedPart3 := "METRICS_PORT: value -12387192873invalid is not a valid integer"
	require.Contains(t, actualLog, expectedPart3)

	expectedPart4 := "failed to validate configuration field KAFKA_SEED_BROKERS: must match ^(|([a-z0-9-]+.[a-z0-9-]+.[a-z]{2,3}"
	require.Contains(t, actualLog, expectedPart4)

	expectedPart5 := "failed to validate configuration field ALERT_TARGET_PREFIX: must match ^((http|https)://|)[a-z0-9-.]+.[a-z]{2,3}/$"
	require.Contains(t, actualLog, expectedPart5)

	expectedPart6 := "failed to validate configuration field ALERT_TARGET_SUFFIX: must match ^@[a-z0-9-]+.[a-z]{2,3}$"
	require.Contains(t, actualLog, expectedPart6)

	expectedPart7 := "failed to validate configuration field VAULT_ENABLED: value what is not a valid boolean value"
	require.Contains(t, actualLog, expectedPart7)

	expectedPart8 := "failed to validate configuration field VAULT_SECRETS_CONFIG: invalid character '}' after top-level value"
	require.Contains(t, actualLog, expectedPart8)
}

func TestAccessors(t *testing.T) {
	docs.Description("the config accessors return the correct values")

	cut, err := tstSetupCutAndLogRecorder(t, "valid-config-unique.yaml")

	require.Nil(t, err)

	actualLog := goauzerolog.RecordedLogForTesting.String()
	require.Equal(t, "", actualLog)

	require.Equal(t, true, cut.PlainLogging())

	require.Equal(t, "some-basic-auth-username", config.Custom(cut).BasicAuthUsername())
	require.Equal(t, "some-basic-auth-password", config.Custom(cut).BasicAuthPassword())
	require.Equal(t, "some-bitbucket-username", config.Custom(cut).BitbucketUsername())
	require.Equal(t, "some-bitbucket-password", config.Custom(cut).BitbucketPassword())
	require.Equal(t, "username", config.Custom(cut).BitbucketReviewerFallback())
	require.Equal(t, "Body, Some", config.Custom(cut).GitCommitterName())
	require.Equal(t, "somebody@somewhere.com", config.Custom(cut).GitCommitterEmail())
	require.Equal(t, "some-kafka-username", config.Custom(cut).KafkaUsername())
	require.Equal(t, "some-kafka-password", config.Custom(cut).KafkaPassword())
	require.Equal(t, "some-kafka-topic", config.Custom(cut).KafkaTopic())
	require.Equal(t, "first-kafka-broker.domain.com:9092,second-kafka-broker.domain.com:9092", config.Custom(cut).KafkaSeedBrokers())
	require.Equal(t, "http://keyset", config.Custom(cut).AuthOidcKeySetUrl())
	require.Equal(t, "some-audience", config.Custom(cut).AuthOidcTokenAudience())
	require.Equal(t, "admin", config.Custom(cut).AuthGroupWrite())
	require.Equal(t, "http://metadata", config.Custom(cut).MetadataRepoUrl())
	require.Equal(t, "git://metadata", config.Custom(cut).SSHMetadataRepositoryUrl())
	require.Equal(t, "5", config.Custom(cut).UpdateJobIntervalCronPart())
	require.Equal(t, uint16(30), config.Custom(cut).UpdateJobTimeoutSeconds())
	require.Equal(t, "https://some-domain.com/", config.Custom(cut).AlertTargetPrefix())
	require.Equal(t, "@some-domain.com", config.Custom(cut).AlertTargetSuffix())
	require.EqualValues(t, []string{"someguy"}, config.Custom(cut).AdditionalPromoters())
	require.EqualValues(t, []string{"add-my-promoters-to-every-service", "also-add-my-promoters"}, config.Custom(cut).AdditionalPromotersFromOwners())
	require.Equal(t, "[a-z][0-1]+", config.Custom(cut).OwnerAliasPermittedRegex().String())
	require.Equal(t, "[a-z][0-2]+", config.Custom(cut).OwnerAliasProhibitedRegex().String())
	require.Equal(t, uint16(1), config.Custom(cut).OwnerAliasMaxLength())
	require.Equal(t, "[a-z][0-3]+", config.Custom(cut).OwnerFilterAliasRegex().String())
	require.Equal(t, "[a-z][0-4]+", config.Custom(cut).ServiceNamePermittedRegex().String())
	require.Equal(t, "[a-z][0-5]+", config.Custom(cut).ServiceNameProhibitedRegex().String())
	require.Equal(t, uint16(2), config.Custom(cut).ServiceNameMaxLength())
	require.Equal(t, "[a-z][0-6]+", config.Custom(cut).RepositoryNamePermittedRegex().String())
	require.Equal(t, "[a-z][0-7]+", config.Custom(cut).RepositoryNameProhibitedRegex().String())
	require.Equal(t, uint16(3), config.Custom(cut).RepositoryNameMaxLength())
	require.Equal(t, ";", config.Custom(cut).RepositoryKeySeparator())
	require.Equal(t, []string{"some-type", "some-other-type"}, config.Custom(cut).RepositoryTypes())
}
