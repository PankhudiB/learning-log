## Kusto

KQL

Played around with the publically available sample data.

Common Kusto query constructs :
https://learn.microsoft.com/en-us/azure/data-explorer/kusto/query/tutorials/learn-common-operators

It seems very similar to SQL

clauses I tried out today : 
- where 
- project --> to select specific columns
- extend --> for derived columns
- sort by XYZ coln
- top 5 by XYZ coln
- distinct XYZ coln
- The MOST fascinating of all ==> 	
	- mapping values from 1 set to another --> using let + combination of query access
	```
	let sourceMapping = dynamic({
          "Emergency Manager" : "Public",     
          "Utility Company" : "Private"
	});
	StormEvents
	| where Source == 'Emergency Manager' or Source == 'Utility Company'
	| project EventId, Source, FriendlyName = sourceMapping[Source];
	
	```

	Some more sample queries I ran :

	```
	StormEvents
	| where State  == 'TEXAS' and EventType  == 'Flood'
	| extend Duration = EndTime - StartTime

	StormEvents
	| project  StartTime,EndTime, Duration = EndTime - StartTime, DamageProperty
	| top 5 by  DamageProperty

	StormEvents
	| sort  by  DamageProperty 

	StormEvents
	| distinct EventType

	StormEvents
	| take 5 
	| project  State, EventType, DamageProperty
	```

Time based filtering -- between clause 

```
StormEvents
| where StartTime between (datetime(2007-01-01) .. datetime(2007-12-31))
```