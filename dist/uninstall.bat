@echo off
:: This needs to be run with an Administrator command prompt
:: See https://technet.microsoft.com/en-us/library/cc947813(v=ws.10).aspx for more information

net stop nlog-metrics
nssm remove nlog-metrics
