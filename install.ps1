$ErrorActionPreference = "Stop"

Write-Host "🥭 Installing Tamarind..." -ForegroundColor Cyan

# Detect Architecture
$Arch = $ENV:PROCESSOR_ARCHITECTURE
if ($Arch -eq "AMD64") {
    $ArchName = "amd64"
} else {
    Write-Host "Unsupported architecture: $Arch. Only 64-bit (AMD64) Windows is currently supported." -ForegroundColor Red
    exit 1
}

$BinaryName = "tamarind-windows-${ArchName}.exe"
$DownloadUrl = "https://github.com/rsantiago/tamarind/releases/download/latest/${BinaryName}"
$ExePath = Join-Path $PWD "tamarind.exe"

Write-Host "⬇️  Downloading ${BinaryName}..."
Invoke-WebRequest -Uri $DownloadUrl -OutFile $ExePath

Write-Host "✅ Tamarind installed successfully in the current directory." -ForegroundColor Green
Write-Host ""
Write-Host "🚀 Run '.\tamarind.exe quickstart' to initialize and boot your server." -ForegroundColor Cyan
