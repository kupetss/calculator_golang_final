# Калькулятор выражений с REST API

Как правильно пользоваться калькулятором

1 Скачайте calculator_golang_final_main.zip

2 Распакуйте архив


3 Откройте папку с проектом в PowerShell и запустите сервер
```powershell
.\calculator.exe server
```

4 Откройте второе окно PowerShell и зарегистрируйте нового пользователя
```powershell
Invoke-RestMethod -Uri "http://localhost:8080/register" `
  -Method Post `
  -Headers @{"Content-Type"="application/json"} `
  -Body '{"username":"123123", "password":"123456"}'
```

5 Выполните вход
```powershell
$loginResponse = Invoke-RestMethod -Uri "http://localhost:8080/login" `
  -Method Post `
  -Headers @{"Content-Type"="application/json"} `
  -Body '{"username":"123123", "password":"123456"}'

$token = $loginResponse.token
```

6 Выполните несколько вычислений:
```powershell
Invoke-RestMethod -Uri "http://localhost:8080/calculate" `
  -Method Post `
  -Headers @{
    "Content-Type"="application/json"
    "Authorization"="Bearer $token"
  } `
  -Body '{"expression":"2 + 2 * 2"}'

Invoke-RestMethod -Uri "http://localhost:8080/calculate" `
  -Method Post `
  -Headers @{
    "Content-Type"="application/json"
    "Authorization"="Bearer $token"
  } `
  -Body '{"expression":"5 * (3 + 1)"}'
```

7 Получите все задачи пользователя:
```powershell
$tasks = Invoke-RestMethod -Uri "http://localhost:8080/tasks" `
  -Method Get `
  -Headers @{"Authorization"="Bearer $token"}

$tasks | ForEach-Object {
    [PSCustomObject]@{
        ID = $_.id
        Выражение = $_.expression
        Статус = $_.status
        Результат = if($_.result){$_.result}else{"-"}
        Дата = $_.created_at
    }
} | Format-Table -AutoSize
```
