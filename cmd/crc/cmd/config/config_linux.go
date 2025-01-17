package config

import (
	cfg "github.com/code-ready/crc/pkg/crc/config"
)

var (
	// Preflight checks
	SkipCheckVirtEnabled             = cfg.AddSetting("skip-check-virt-enabled", nil, []cfg.ValidationFnType{cfg.ValidateBool}, []cfg.SetFn{cfg.SuccessfullyApplied})
	WarnCheckVirtEnabled             = cfg.AddSetting("warn-check-virt-enabled", nil, []cfg.ValidationFnType{cfg.ValidateBool}, []cfg.SetFn{cfg.SuccessfullyApplied})
	SkipCheckKvmEnabled              = cfg.AddSetting("skip-check-kvm-enabled", nil, []cfg.ValidationFnType{cfg.ValidateBool}, []cfg.SetFn{cfg.SuccessfullyApplied})
	WarnCheckKvmEnabled              = cfg.AddSetting("warn-check-kvm-enabled", nil, []cfg.ValidationFnType{cfg.ValidateBool}, []cfg.SetFn{cfg.SuccessfullyApplied})
	SkipCheckLibvirtInstalled        = cfg.AddSetting("skip-check-libvirt-installed", nil, []cfg.ValidationFnType{cfg.ValidateBool}, []cfg.SetFn{cfg.SuccessfullyApplied})
	WarnCheckLibvirtInstalled        = cfg.AddSetting("warn-check-libvirt-installed", nil, []cfg.ValidationFnType{cfg.ValidateBool}, []cfg.SetFn{cfg.SuccessfullyApplied})
	SkipCheckLibvirtEnabled          = cfg.AddSetting("skip-check-libvirt-enabled", nil, []cfg.ValidationFnType{cfg.ValidateBool}, []cfg.SetFn{cfg.SuccessfullyApplied})
	WarnCheckLibvirtEnabled          = cfg.AddSetting("warn-check-libvirt-enabled", nil, []cfg.ValidationFnType{cfg.ValidateBool}, []cfg.SetFn{cfg.SuccessfullyApplied})
	SkipCheckLibvirtRunning          = cfg.AddSetting("skip-check-libvirt-running", nil, []cfg.ValidationFnType{cfg.ValidateBool}, []cfg.SetFn{cfg.SuccessfullyApplied})
	WarnCheckLibvirtVersionCheck     = cfg.AddSetting("warn-check-libvirt-version", nil, []cfg.ValidationFnType{cfg.ValidateBool}, []cfg.SetFn{cfg.SuccessfullyApplied})
	SkipCheckLibvirtVersionCheck     = cfg.AddSetting("skip-check-libvirt-version", nil, []cfg.ValidationFnType{cfg.ValidateBool}, []cfg.SetFn{cfg.SuccessfullyApplied})
	WarnCheckLibvirtRunning          = cfg.AddSetting("warn-check-libvirt-running", nil, []cfg.ValidationFnType{cfg.ValidateBool}, []cfg.SetFn{cfg.SuccessfullyApplied})
	SkipCheckUserInLibvirtGroup      = cfg.AddSetting("skip-check-user-in-libvirt-group", nil, []cfg.ValidationFnType{cfg.ValidateBool}, []cfg.SetFn{cfg.SuccessfullyApplied})
	WarnCheckUserInLibvirtGroup      = cfg.AddSetting("warn-check-user-in-libvirt-group", nil, []cfg.ValidationFnType{cfg.ValidateBool}, []cfg.SetFn{cfg.SuccessfullyApplied})
	SkipCheckLibvirtDriver           = cfg.AddSetting("skip-check-libvirt-driver", nil, []cfg.ValidationFnType{cfg.ValidateBool}, []cfg.SetFn{cfg.SuccessfullyApplied})
	WarnCheckLibvirtDriver           = cfg.AddSetting("warn-check-libvirt-driver", nil, []cfg.ValidationFnType{cfg.ValidateBool}, []cfg.SetFn{cfg.SuccessfullyApplied})
	SkipCheckCrcNetwork              = cfg.AddSetting("skip-check-crc-network", nil, []cfg.ValidationFnType{cfg.ValidateBool}, []cfg.SetFn{cfg.SuccessfullyApplied})
	WarnCheckCrcNetwork              = cfg.AddSetting("warn-check-crc-network", nil, []cfg.ValidationFnType{cfg.ValidateBool}, []cfg.SetFn{cfg.SuccessfullyApplied})
	SkipCheckCrcNetworkActive        = cfg.AddSetting("skip-check-crc-network-active", nil, []cfg.ValidationFnType{cfg.ValidateBool}, []cfg.SetFn{cfg.SuccessfullyApplied})
	WarnCheckCrcNetworkActive        = cfg.AddSetting("warn-check-crc-network-active", nil, []cfg.ValidationFnType{cfg.ValidateBool}, []cfg.SetFn{cfg.SuccessfullyApplied})
	SkipCheckCrcDnsmasqFile          = cfg.AddSetting("skip-check-crc-dnsmasq-file", nil, []cfg.ValidationFnType{cfg.ValidateBool}, []cfg.SetFn{cfg.SuccessfullyApplied})
	WarnCheckCrcDnsmasqFile          = cfg.AddSetting("warn-check-crc-dnsmasq-file", nil, []cfg.ValidationFnType{cfg.ValidateBool}, []cfg.SetFn{cfg.SuccessfullyApplied})
	SkipCheckCrcNetworkManagerConfig = cfg.AddSetting("skip-check-network-manager-config", nil, []cfg.ValidationFnType{cfg.ValidateBool}, []cfg.SetFn{cfg.SuccessfullyApplied})
	WarnCheckCrcNetworkManagerConfig = cfg.AddSetting("warn-check-network-manager-config", nil, []cfg.ValidationFnType{cfg.ValidateBool}, []cfg.SetFn{cfg.SuccessfullyApplied})
	WarnCheckNetworkManagerInstalled = cfg.AddSetting("warn-check-network-manager-installed", nil, []cfg.ValidationFnType{cfg.ValidateBool}, []cfg.SetFn{cfg.SuccessfullyApplied})
	SkipCheckNetworkManagerInstalled = cfg.AddSetting("skip-check-network-manager-installed", nil, []cfg.ValidationFnType{cfg.ValidateBool}, []cfg.SetFn{cfg.SuccessfullyApplied})
	WarnCheckNetworkManagerRunning   = cfg.AddSetting("warn-check-network-manager-running", nil, []cfg.ValidationFnType{cfg.ValidateBool}, []cfg.SetFn{cfg.SuccessfullyApplied})
	SkipCheckNetworkManagerRunning   = cfg.AddSetting("skip-check-network-manager-running", nil, []cfg.ValidationFnType{cfg.ValidateBool}, []cfg.SetFn{cfg.SuccessfullyApplied})
)
