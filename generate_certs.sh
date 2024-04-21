# Certificate Authority (CA)
openssl req -new -x509 \
  -sha256 \
  -subj "/CN=ca.localhost/C=DE" \
  -days 3650 \
  -keyout ca-key.pem \
  -out ca-cert.pem

# Server certificate
# Create a key
openssl genrsa -out server/cert/key.pem 4096
# Create a certificate signing request
openssl req -new \
  -subj "/CN=localhost/C=DE" \
  -addext "subjectAltName=DNS:localhost" \
  -key server/cert/key.pem \
  -out server/cert/csr.pem
# Sign it with the CA
openssl x509 -req \
  -extfile <(printf "subjectAltName=DNS:localhost") \
  -days 3650 \
  -in server/cert/csr.pem \
  -CA ca-cert.pem \
  -CAkey ca-key.pem -CAcreateserial \
  -out server/cert/cert.pem

# Client certificate
# Create a key
openssl genrsa -out client/cert/key.pem 4096
# Create a certificate signing request
openssl req -new \
  -subj "/CN=localhost/C=DE" \
  -key client/cert/key.pem \
  -out client/cert/csr.pem
# Sign it with the CA
openssl x509 -req \
  -days 3650 \
  -in client/cert/csr.pem \
  -CA ca-cert.pem \
  -CAkey ca-key.pem -CAcreateserial \
  -out client/cert/cert.pem

