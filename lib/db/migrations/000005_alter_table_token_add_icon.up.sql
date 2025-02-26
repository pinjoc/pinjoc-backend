ALTER TABLE token ADD COLUMN icon TEXT;

UPDATE token SET icon = 'https://etherscan.io/token/images/wrappedbtc_ofc_32.svg' WHERE id = 1;
UPDATE token SET icon = 'https://etherscan.io/token/images/usdc_ofc_32.svg' WHERE id = 2;
UPDATE token SET icon = 'https://etherscan.io/token/images/weth_28.png?v=2' WHERE id = 3;
UPDATE token SET icon = 'https://etherscan.io/token/images/dairplce_32.svg' WHERE id = 4;
UPDATE token SET icon = 'https://etherscan.io/token/images/aaverplce_32.svg' WHERE id = 5;
