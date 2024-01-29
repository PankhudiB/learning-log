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