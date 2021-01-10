# How to generate a message signature

> Source: https://docs.firstdata.com/org/gateway/node/394
 

## Purpose
To ensure data integrity, prevent replay attacks, and eliminate stale requests.

## Details
- Signature Algorithm: SHA256 HMAC
- Signature Encoding: Base64
- Signed with: Developer App Secret Key; provided to merchant when boarded

The message data for the signature is the following items concatenated:

- API_KEY: Developer App Api Key; provided to merchant when boarded
- CLIENT_REQUEST_ID: Unique request id, 128-bit UUIDv4 format recommended
- TIMESTAMP: Epoch timestamp in milliseconds
- PAYLOAD: The request body if applicable

The ClientRequestID is a randomly generated number that is unique for each request. It is used as nonce and validated against all ClientRequestIDs received by First Data within a predetermined timeframe (five minutes is the default) to prevent replay attacks. First Data uses the timestamp of the request to validate against stale requests. Any request older than the specified duration is rejected. 

## Code Examples
 
### Javascript

```js
let msg = API_KEY + CLIENT_REQUEST_ID + TIMESTAMP + JSON.stringify(PAYLOAD);

const hmac = CryptoJS.algo.HMAC.create(CryptoJS.algo.SHA256, API_SECRET);
hmac.update(msg);
const messageSignature = base64.encode(hmac.finalize().toString());
```
 
### PHP

```php
$msg = $API_KEY . $CLIENT_REQUEST_ID . $TIMESTAMP . $JSON_SERIALIZED_PAYLOAD;
$hmac = hash_hmac('sha256', $msg, $API_SECRET);
$messageSignature = base64_encode($hmac);
```
 
### Java

```java
import org.apache.commons.codec.binary.Base64;
import org.apache.commons.codec.binary.Hex;
import org.apache.commons.codec.digest.HmacAlgorithms;
import org.apache.commons.codec.digest.HmacUtils;

public class Example {
    public static void main(String[] args) {
        final HmacUtils hmacHelper = new HmacUtils(HmacAlgorithms.HMAC_SHA_256, API_SECRET);
        final Hex hexHelper = new Hex();

        final String msg = API_KEY + CLIENT_REQUEST_ID + TIMESTAMP + JSON_SERIALIZED_PAYLOAD;
        final byte[] raw = hmacHelper.hmac(msg);
        final byte[] hex = hexHelper.encode(raw);
        final String messageSignature = Base64.encodeBase64String(hex);
    }
}
```
