# Калькулятор выражений с REST API

Как правильно пользоваться калькулятором

1 Сначала зарегистрируйте нового пользователя
```powershell
Invoke-RestMethod -Uri "http://localhost:8080/register" `
  -Method Post `
  -Headers @{"Content-Type"="application/json"} `
  -Body '{"username":"123123", "password":"123456"}'
```

2 Выполните вход
```powershell
$loginResponse = Invoke-RestMethod -Uri "http://localhost:8080/login" `
  -Method Post `
  -Headers @{"Content-Type"="application/json"} `
  -Body '{"username":"123123", "password":"123456"}'

$token = $loginResponse.token
```

3 Выполните несколько вычислений:
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

4 Получите все задачи пользователя:
```powershell
$tasks = Invoke-RestMethod -Uri "http://localhost:8080/tasks" `
  -Method Get `
  -Headers @{"Authorization"="Bearer $token"}

# Красивое отображение
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