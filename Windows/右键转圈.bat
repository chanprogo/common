regsvr32 /u /s igfxpph.dll
reg delete HKEY_CLASSES_ROOTDirectoryBackgroundshellexContextMenuHandlers /f
reg add HKEY_CLASSES_ROOTDirectoryBackgroundshellexContextMenuHandlersnew /ve /d {D969A300-E7FF-11d0-A93B- 00A0C90F2719}
reg delete HKEY_LOCAL_MACHINESOFTWAREMicrosoftWindowsCurrentVersionRun /v HotKeysCmds /f
reg delete HKEY_LOCAL_MACHINESOFTWAREMicrosoftWindowsCurrentVersionRun /v IgfxTray /f