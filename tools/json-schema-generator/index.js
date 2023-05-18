import { parseArgs } from 'node:util';
import { readFile, writeFile } from 'node:fs/promises';
import { Parser } from '@asyncapi/parser';
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

	const apiContent = await readFile(apiFile, 'utf8');

	const { document } = await new Parser().parse(apiContent);

	const {
		_json: { channels }
	} = document;

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
