# stew

stew means __Stack Trace Error Wrap__

## Before using stew

Your error message is like

```
dial tcp [::1]:3306: connect: connection refused
```

What the hell?! Where is the cause line?

## After using stew

Just replace `return err` to `return stew.Wrap(err)` then

```
"dial tcp [::1]:3306: connect: connection refused FullPath/repositories.go:106 github.com/momotaro98/mixlunch-service-api/userscheduleservice.(*realUserScheduleQueryRepository).QueryUserSchedulesWhereTimeRange
FullPath/domain.go:129 github.com/momotaro98/mixlunch-service-api/userscheduleservice.(*realUserScheduleServer).GetUserSchedulesByTimeRange
```
