param (
    [string]$Path = ".",
    [string[]]$ExcludeFolder = @("node_modules", "bin", "obj", ".git", ".vs"),
    [string[]]$IncludeExtension = @(".go", ".cs", ".js", ".ts", ".html", ".css", ".sql", ".yaml", ".yml", ".json", ".sh", ".ps1")
)

$files = Get-ChildItem -Path $Path -Recurse -File | 
    Where-Object { 
        $includePath = $true
        foreach ($exclude in $ExcludeFolder) {
            if ($_.FullName -like "*\$exclude\*") {
                $includePath = $false
                break
            }
        }
        $includePath -and ($IncludeExtension -contains $_.Extension)
    }

$totalLines = 0
$extensionStats = @{}

foreach ($file in $files) {
    $lines = (Get-Content $file.FullName | Measure-Object -Line).Lines
    $totalLines += $lines
    
    if (!$extensionStats.ContainsKey($file.Extension)) {
        $extensionStats[$file.Extension] = @{
            FileCount = 0
            LineCount = 0
        }
    }
    
    $extensionStats[$file.Extension].FileCount++
    $extensionStats[$file.Extension].LineCount += $lines
}

Write-Output "`nFile Extension Statistics:"
Write-Output "------------------------"
foreach ($ext in $extensionStats.Keys | Sort-Object) {
    Write-Output ("{0,-8} Files: {1,6:N0} Lines: {2,8:N0}" -f 
        $ext,
        $extensionStats[$ext].FileCount,
        $extensionStats[$ext].LineCount)
}

Write-Output "`nSummary:"
Write-Output "--------"
Write-Output "Total Files: $($files.Count)"
Write-Output "Total Lines: $totalLines"