package profile

import (
	"fmt"
	"os"
	"strings"

	"github.com/bigkevmcd/go-configparser"
	"github.com/golinuxcloudnative/aws-export-sso-profile/internal/domain"
	"go.uber.org/zap"
)

type Profile struct {
	log           *zap.Logger
	AwsConfigFile string
}

func NewProfile(awsConfigFile string) *Profile {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	if awsConfigFile == "" {
		// get user's home dir
		homeDir, errHome := os.UserHomeDir()
		if errHome != nil {
			logger.Fatal("could not get user's home dir", zap.Error(errHome))
			return nil
		}
		awsConfigFile = fmt.Sprintf("%s/%s", homeDir, ".aws/config")
	}

	logger.Debug("current aws config file", zap.String("awsConfigFile", awsConfigFile))

	return &Profile{
		log:           logger,
		AwsConfigFile: awsConfigFile,
	}

}

func (p *Profile) Profiles() ([]*domain.Profile, error) {
	configParsed, err := p.readConfig()
	if err != nil {
		return nil, err
	}
	profiles, err := p.listProfiles(configParsed)
	if err != nil {
		return nil, err
	}
	return profiles, nil
}

func (p *Profile) readConfig() (*configparser.ConfigParser, error) {

	configParsed, errCfgParser := configparser.NewConfigParserFromFile(p.AwsConfigFile)
	if errCfgParser != nil {
		return nil, fmt.Errorf("could not parse file, %w", errCfgParser)
	}

	p.log.Debug("number of profiles", zap.Int("profiles", len(configParsed.Sections())))

	return configParsed, nil
}

func (p *Profile) listProfiles(configParsed *configparser.ConfigParser) ([]*domain.Profile, error) {
	profiles := []*domain.Profile{}
	for _, name := range configParsed.Sections() {
		config, _ := configParsed.Items(name)

		p.log.Debug("name of profile", zap.String("profile", name), zap.Any("props", config))

		// convert one profile of config file to profile struct
		profile := p.configToProfile(name, config)

		profiles = append(profiles, profile)
	}

	return profiles, nil
}

func (p *Profile) configToProfile(name string, config map[string]string) *domain.Profile {
	p.log.Debug("name of profile", zap.String("profile", name), zap.Any("props", config))
	opts := &domain.ProfileOptions{}
	if val, ok := config["region"]; ok {
		p.log.Debug("region value", zap.String("profile", name), zap.String("region", val))
		opts.Region = val
	}
	if val, ok := config["sso_account_id"]; ok {
		p.log.Debug("sso_account_id value", zap.String("profile", name), zap.String("sso_account_id", val))
		opts.Sso_account_id = val
	}
	if val, ok := config["sso_region"]; ok {
		p.log.Debug("sso_region value", zap.String("profile", name), zap.String("sso_region", val))
		opts.Sso_region = val
	}
	if val, ok := config["sso_role_name"]; ok {
		p.log.Debug("sso_role_name value", zap.String("profile", name), zap.String("sso_role_name", val))
		opts.Sso_role_name = val
	}
	if val, ok := config["sso_start_url"]; ok {
		p.log.Debug("sso_start_url value", zap.String("profile", name), zap.String("sso_start_url", val))
		opts.Sso_start_url = val
	}

	// remove profile
	profileName := strings.Replace(name, "profile ", "", 1)
	// create profile
	profile := domain.NewProfile(profileName, opts)

	p.log.Debug("new profile", zap.Any("profile", profile))

	return profile
}
