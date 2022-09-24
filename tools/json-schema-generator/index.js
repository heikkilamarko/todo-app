import { readFile, writeFile } from 'fs/promises';
import { Command } from 'commander';
import parser from '@asyncapi/parser';
import chalk from 'chalk';
import YAML from 'yaml';

try {
	const { config } = parseFlags();

	const { services } = YAML.parse(await readFile(config, 'utf8'));

	for (const service of services) {
		logDivider();
		await generate(service);
	}
	logDivider();
} catch (err) {
	logFatal('Error generating json schemas.', err);
}

function parseFlags() {
	const program = new Command();
	program.requiredOption('-c, --config <file>', 'config file');
	program.parse();
	return program.opts();
}

async function generate({ name, apiFile, outputDir }) {
	console.log(chalk.yellowBright('service'));
	console.log(chalk.cyanBright(`    ${name}`));
	console.log(chalk.yellowBright('input'));
	console.log(chalk.cyanBright(`    ${apiFile}`));
	console.log(chalk.yellowBright('output'));

	const apiContent = await readFile(apiFile, 'utf8');

	let {
		_json: { channels }
	} = await parser.parse(apiContent);

	const messages = Object.entries(channels)
		.filter(([_, c]) => c.publish)
		.map(([name, c]) => ({
			file: `${outputDir}/${name}.json`,
			schema: JSON.stringify(c.publish.message.payload, null, 2)
		}));

	for (const m of messages) {
		await writeFile(m.file, m.schema);
		console.log(chalk.green(`    ${m.file}`));
	}
}

function logDivider() {
	console.log();
	console.log(chalk.yellowBright(`----------------------------------------`));
	console.log();
}

function logFatal(message, err) {
	console.log(chalk.redBright(message));
	if (err) {
		console.log(chalk.red(err));
	}
	process.exit(1);
}
