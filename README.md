## gofire

Apache Geode Golang client. Limited to the REST API, right now.

Should be pretty straightforward, easy to use, simple.

Proof:

```go
// insecure
geode, _ := NewClient("https://1.2.3.4:8080", true)
geode.Region = "exampleRegion"

type omgAType struct { wutField int }

lolwut := omgAType{wutField: 1}
var holdMeDaddy omgAType
geode.Put("deyTookUrData", lolwut)
geode.Get("deyTookUrData", &holdMeDaddy)
```

## Sidebar

Geode uses this really wonderful concept called dynamic JSON for the `/gemfire-api/v1/{region}` API. Example:

```json
{
  "orders" : [ {
     "purchaseOrderNo" : 1112,
     "customerId" : 102,
     "description" :  "Purchase order for  company - B",
     "orderDate" :  "02/10/2014",
     "deliveryDate" :  "02/20/2014",
     "contact" :  "John Doe",
     "email" :  "John.Doe@example.com",
     "phone" :  "01-2048096",
     "items" : [ {
       "itemNo" : 1,
       "description" :  "Product-AAAA",
       "quantity" : 10,
       "unitPrice" : 20.0,
       "totalPrice" : 200.0
    }, {
       "itemNo" : 2,
       "description" :  "Product-BBB",
       "quantity" : 15,
       "unitPrice" : 10.0,
       "totalPrice" : 150.0
    } ],
     "totalPrice" : 350.0
  }, {
     "purchaseOrderNo" : 111,
     "customerId" : 101,
     "description" :  "Purchase order for  company - A",
     "orderDate" :  "01/10/2014",
     "deliveryDate" :  "01/20/2014",
     "contact" :  "Jane Doe",
     "email" :  "Jane.Doe@example.com",
     "phone" :  "020-2048096",
     "items" : [ {
       "itemNo" : 1,
       "description" :  "Product-1",
       "quantity" : 5,
       "unitPrice" : 10.0,
       "totalPrice" : 50.0
    }, {
       "itemNo" : 1,
       "description" :  "Product-2",
       "quantity" : 10,
       "unitPrice" : 15.5,
       "totalPrice" : 155.0
    } ],
     "totalPrice" : 205.0
  } ]
}
```

Every single field in that JSON is dynamic. The `orders` field changes depending on what region you are using, and all the data in the array can be of any structure. So. Without enforcing a lot of nonsense, I've implemented the API basically, but for your own sanity, I don't recommend using it with this client. I tried, realised I don't hate myself, and I stopped. There is a way to do it, but it will take a lot of rum and pain. You buy the rum.