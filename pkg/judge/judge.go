package judge

type Judge struct {
	BasePath  string
	Languages []*Language
	Rules     map[string]*Rule
}

func InitJudge(judgeConfigPath string) (*Judge, error) {
	judge := &Judge{
		BasePath: judgeConfigPath + "/",
	}

	if err := judge.LoadLanguages(); err != nil {
		return nil, err
	}
	if err := judge.LoadRules(); err != nil {
		return nil, err
	}

	return judge, nil
}
