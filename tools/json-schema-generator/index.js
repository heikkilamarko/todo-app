import { parseArgs } from 'node:util';
import { readFile, writeFile } from 'node:fs/promises';
import { Parser, fromFile } from '@asyncapi/parser';
import chalk from 'chalk';
import YAML from 'yaml';

try {
	const {
		values: { config }
	} = parseArgs({
		options: {
			config: {
				type: 'string',
				short: 'c',
				default: 'config.yaml'
			}
		}
	});

	const { services } = YAML.parse(await readFile(config, 'utf8'));

	for (const service of services) {
		logDivider();
		await generate(service);
	}
	logDivider();
} catch (err) {
	logFatal('Error generating json schemas.', err);
}

async function generate({ name, apiFile, outputDir }) {
	console.log(chalk.yellowBright('service'));
	console.log(chalk.cyanBright(`    ${name}`));
	console.log(chalk.yellowBright('input'));
	console.log(chalk.cyanBright(`    ${apiFile}`));
	console.log(chalk.yellowBright('output'));

	const parser = new Parser();

	const {
		document: {
			_json: { operations }
		}
	} = await fromFile(parser, apiFile).parse();

	const messages = Object.entries(operations).map(([_, o]) => ({
		file: `${outputDir}/${o.channel.address}.json`,
		schema: JSON.stringify(o.messages[0].payload, null, 2)
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
