package netmgr

func (nm *networkManager) StateChanged(state chan<- StateEnum) error {
	return nm.USignal(NetworkManagerInterface, "StateChanged", state, nil)
}

func StateChanged(state chan<- StateEnum) error {
	nm, err := System()
	if err != nil {
		return err
	}
	return nm.StateChanged(state)
}
