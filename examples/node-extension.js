#!/usr/bin/env node

const args = process.argv.slice(2);

console.log("🟢 Node.js Extension Example");
console.log("================================");
console.log(`Node version: ${process.version}`);
console.log(`Platform: ${process.platform}`);
console.log(`Arguments: ${JSON.stringify(args)}`);
console.log("================================");
console.log("\n✨ This extension runs with Node.js!");
