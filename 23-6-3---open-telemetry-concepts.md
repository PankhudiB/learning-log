## OpenTelemetry Concepts 
(Understood through reading docs)

I have used open-census and open-tracing in my projects. 
I was curious and have been wanting to give open-telemetry a try.

OTEL is basically an initiative to standardise observability specs, APIs, SDK integrations - so that there is no vendor lock-in.     

OTEL has concepts of `Signals` - `which felt like umbrella term for - traces, metrics, logs and baggage.` 

Now, I am already aware of the concepts & have hands-on on traces, metrics.
Traces - while I was implementing distributed-tracing solution across microservices - Had used open-census + Jaeger
Metrics - while I had integrated Prometheus metrics. 
Logs - I believe should be more or less same. 

`Baggage` was something that I found new.

Upon reading specs and docs,
I understood the use-case is when you want (& are comfortable) to pass something custom key-value across spans and traces.
`Eg : customer-id, client-IP.` 
So, this `propagation happens through HTTP headers`

If you are okay exposing that info - go ahead ! leverage baggage ! 

Although at very first glance - it looked like this baggage is very similar to Tags or more formally "Span attributes"

But apparently - otel docs addresses that and says "they differ from span attributes"

span attributes -> passed across spans 

baggage -> you have to explicitly extract and forward

```
   var accountId = Baggage.GetBaggage("AccountId");
    Activity.Current?.SetTag("AccountId", accountId);
```

In another place they have mentioned :

`Baggage` is intended for indexing observability events in one service with attributes provided by a prior service in the same transaction.
This helps to establish a causal relationship between these events.

`I am still not sure why did they have to distinguish ?`