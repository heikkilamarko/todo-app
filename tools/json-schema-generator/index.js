import { readFile, writeFile } from "fs/promises";
import parser from "@asyncapi/parser";
import chalk from "chalk";
import YAML from "yaml";

const configFile = process.argv.slice(2)[0];

if (!configFile) {
  console.log(
    chalk.red("error: You must pass a config file (yaml) as the only argument.")
  );
  process.exit(1);
}

let config = await readFile(configFile, "utf8");

config = YAML.parse(config);

for (const service of config.services) {
  logDivider();
  await generate(service);
}

logDivider();

async function generate({ name, apiFile, outputDir }) {
  console.log(chalk.yellow("service"));
  console.log(chalk.cyan(`    ${name}`));
  console.log(chalk.yellow("input"));
  console.log(chalk.cyan(`    ${apiFile}`));
  console.log(chalk.yellow("output"));

  const apiContent = await readFile(apiFile, "utf8");

  let {
    _json: { channels },
  } = await parser.parse(apiContent);

  const messages = Object.entries(channels)
    .filter(([_, c]) => c.publish)
    .map(([name, c]) => ({
      file: `${outputDir}/${name}.json`,
      schema: JSON.stringify(c.publish.message.payload, null, 2),
    }));

  for (const m of messages) {
    await writeFile(m.file, m.schema);
    console.log(chalk.green(`    ${m.file}`));
  }
}

function logDivider() {
  console.log();
  console.log(chalk.yellow(`----------------------------------------`));
  console.log();
}
