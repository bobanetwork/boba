# Alchemy Pay URL Generator Service

A simple serverless service to generate signed URLs for Alchemy Pay integration.

## Prerequisites

- Node.js (v14 or later)
- Python 3.9
- AWS CLI configured with appropriate credentials
- Serverless Framework

## Quick Start

1. **Install Dependencies**
```bash
# Install Serverless Framework globally
npm install -g serverless

# Install project dependencies
pnpm install

# Install Python dependencies
pip install -r requirements.txt
```

2. **Configure Environment**
```bash
# Copy example environment files
cp env-dev.example.yml env-dev.yml
cp env-mainnet.example.yml env-mainnet.yml

# Edit the files with your credentials
```

3. **Local Testing**
```bash
# Start local serverless offline
pnpm run dev

# Test URL generation
curl -X POST http://localhost:3000/dev/generate_alchemypay_url \
  -H "Content-Type: application/json" \
  -d '{
    "crypto": "USDT",
    "fiatAmount": "15",
    "fiat": "USD",
    "merchantOrderNo": "test123",
    "network": "BSC"
  }'

# Test the routescan api invocation
curl -X POST http://localhost:3000/dev/fetch_routescan_l2_transaction \
  -H "Content-Type: application/json" \
  -d '{
    "chainId": "56288",
    "address": "0x9703d3B2521F3De2D56831f3df9490cbB1487428",
    "startBlock": "41661017"
  }'

```

4. **Deployment**
```bash
# Deploy to dev
./deploy.sh dev

# Deploy to mainnet (requires MAINNET_SECRET_KEY env variable)
./deploy.sh mainnet
```
