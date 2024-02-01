## Kusto Aggregations

KQL

summarize by --> is like group by
render --> to visaualise the summarized operation

	```
	StormEvents
	| summarize TotalStorms = count() by State
	| render barchart
	```

---

countif(predicate) -> to conditionally count

	```
	StormEvents
	| summarize TotalStorms = countif(DamageCrops > 0) by State
	| top 5 by TotalStorms
	```
This seemed super cool - so a custom field is created -> that uses `countif` -- and the same custom field is used for `top 5`

---

You can summarize by bins -- like buckets - similar to how histogram does it, and further visualise it with `timechart`

```
StormEvents
| summarize EventCount = count() by bin(StartTime, 30d)
| render timechart
```

---

summarize to do min.max,avg

```
StormEvents
| summarize 
    MaxCropDamage=max(DamageCrops),
    MinCropDamage=min(DamageCrops),
    AvgCropDamage=avg(DamageCrops)
    by EventType
| sort by AvgCropDamage
```

-----

Another interesting one --> EXTRACT UNIQUE VALUES with `make_set()`

To get all unique event types that caused death within a state:

```
StormEvents
| where DeathsDirect > 0 or DeathsIndirect > 0
| summarize StormTypesWithDeath = make_set(EventType) by State
| project State, StormTypesWithDeath
```

output would look something like : 
```
STATE | StormTypesWithDeath
State1 :  [X,Y,Z]
State2 :  [A,B,C]
```