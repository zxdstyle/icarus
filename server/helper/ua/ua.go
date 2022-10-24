package ua

import (
	"github.com/mssola/user_agent"
)

const (
	DeviceMobile  = "Mobile"
	DeviceDesktop = "Desktop"
	DeviceOther   = "Other"
)

type UserAgent struct {
	agent *user_agent.UserAgent
}

func New(agent string) *UserAgent {
	ua := &UserAgent{}
	ua.doParseUserAgent(agent)
	return ua
}

func (ua *UserAgent) doParseUserAgent(agent string) {
	ua.agent = user_agent.New(agent)
}

func (ua *UserAgent) Browser() (browser string, version string) {
	return ua.agent.Browser()
}

func (ua *UserAgent) Platform() string {
	return ua.agent.Platform()
}

func (ua *UserAgent) Device() string {
	if ua.agent.Mobile() {
		return DeviceMobile
	}
	return DeviceDesktop
}
