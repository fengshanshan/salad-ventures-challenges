const { ethers } = require("ethers");

const ADDR = "0xc7a9bC274e77c77517A6b0eA0BaB98d8807cA0c1";   // your contract address
const ABI = [{"inputs":[{"internalType":"address","name":"addr","type":"address"},{"internalType":"address[]","name":"tokens","type":"address[]"}],"name":"getBalances","outputs":[{"components":[{"internalType":"address","name":"token","type":"address"},{"internalType":"uint256","name":"balance","type":"uint256"}],"internalType":"struct Utility.TokenBalances[]","name":"","type":"tuple[]"}],"stateMutability":"view","type":"function"}];    // your contract ABI

const ADDRESS = "0xbfDC6603DC5938D9d75b580c92b280183D4db020"; // some wallet address with token balance
const TOKENS = [    // token contract addresses
    "0x3beb68535B2693a92e2fC6d4B0911F637f224846", //PUX
    "0x126E68610A4b07401a9f4Af73c3C66A55b84Fe16",
];

const provider = new ethers.providers.JsonRpcProvider(
    "https://rinkeby.infura.io/v3/816cbf2267454e52ac7633d5429dc5d5"
); // rinkeby

const test = async () => {
    const contract = new ethers.Contract(ADDR, ABI, provider);

    const balances = await contract.getBalances(ADDRESS, TOKENS);

    console.log(balances.toString())

    return balances;
};

test().then(console.log);