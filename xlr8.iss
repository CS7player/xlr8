[Setup]
AppName=xlr8
AppVersion=1.0.0
DefaultDirName={pf}\xlr8
DefaultGroupName=xlr8
OutputBaseFilename=xlr8-installer
Compression=lzma
SolidCompression=yes
ArchitecturesInstallIn64BitMode=x64

; Installer icon
SetupIconFile=xlr8.ico

[Files]
Source: "dist\xlr8.exe"; DestDir: "{app}"; Flags: ignoreversion

[Icons]
Name: "{group}\xlr8"; Filename: "{app}\xlr8.exe"; IconFilename: "{app}\xlr8.exe"
Name: "{commondesktop}\xlr8"; Filename: "{app}\xlr8.exe"; Tasks: desktopicon; IconFilename: "{app}\xlr8.exe"

[Tasks]
Name: desktopicon; Description: "Create a desktop shortcut"; Flags: unchecked
Name: addtopath; Description: "Add xlr8 to system PATH"; Flags: unchecked

[Code]
procedure AddToPath(Dir: string);
var
  Path: string;
begin
  if not RegQueryStringValue(HKEY_CURRENT_USER,
    'Environment', 'Path', Path) then
    Path := '';

  if Pos(Dir, Path) = 0 then
  begin
    if Path <> '' then
      Path := Path + ';' + Dir
    else
      Path := Dir;

    RegWriteStringValue(HKEY_CURRENT_USER,
      'Environment', 'Path', Path);
  end;
end;

procedure CurStepChanged(CurStep: TSetupStep);
begin
  if CurStep = ssPostInstall then
  begin
    if WizardIsTaskSelected('addtopath') then
      AddToPath(ExpandConstant('{app}'));
  end;
end;