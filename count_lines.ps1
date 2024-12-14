# CountLines.ps1 - Recursively count total lines in files within a folder

param (
    [string]$Path = "."
)

$totalLines = Get-ChildItem -Path $Path -Recurse -File | 
    ForEach-Object { 
        (Get-Content $_.FullName | Measure-Object -Line).Lines
    } | Measure-Object -Sum

Write-Output "Total lines in all files: $($totalLines.Sum)"
