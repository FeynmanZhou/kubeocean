package ssh

func CmdExec(host string, user string, port string, password string, force bool, cmd string) error {
	puser := NewUser(user, port, password, force)
	err := SingleRun(host, cmd, puser, force)
	if err != nil {
		return err
	}
	return nil
}

func SingleRun(host string, cmd string, user *CommonUser, force bool) error {
	server := NewCmdServer(host, user.port, user.user, user.psw, "cmd", cmd, force)
	r := server.SRunCmd()
	PrintExecResult(r)
	if r.Err != nil {
		return r.Err
	}
	return nil
}

func CmdExecOut(host string, user string, port string, password string, force bool, cmd string) (string, error) {
	puser := NewUser(user, port, password, force)
	out, err := SingleRunOut(host, cmd, puser, force)
	if err != nil {
		return "", err
	} else {
		return out, nil
	}
}

func SingleRunOut(host string, cmd string, user *CommonUser, force bool) (string, error) {
	server := NewCmdServer(host, user.port, user.user, user.psw, "cmd", cmd, force)
	r := server.SRunCmd()
	PrintExecResult(r)
	if r.Err != nil {
		return "", r.Err
	} else {
		return r.Result, nil
	}

}
