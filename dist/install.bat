@echo off
:: This needs to be run with an Administrator command prompt
:: See https://technet.microsoft.com/en-us/library/cc947813(v=ws.10).aspx for more information

:: Update the path to match your machine:
nssm install nlog-metrics d:\services\nlog-metrics\nlog-metrics_windows_amd64.exe
nssm set nlog-metrics AppParameters start
nssm set nlog-metrics Description NLog metrics collection and logging
nssm set nlog-metrics Start SERVICE_AUTO_START

net start nlog-metrics