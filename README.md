# special-days

This is a package to give dates for mothering Sunday and fathers day (UK).

```
go get github.com/charlieegan3/special-days
```

```
fathersDay, err :=  fathersDay.FathersDay("uk", 2021) // => time.Time, err
```

```
fathersDay, err := motheringsunday.MotheringSunday(2021) // => time.Time, err
```
