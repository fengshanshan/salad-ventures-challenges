const Uti = artifacts.require("Utility");

module.exports = async function (deployer) {
  await deployer.deploy(Uti);
  const utility = await Uti.deployed()
  console.log(utility.address)
};
