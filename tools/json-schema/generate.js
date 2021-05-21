import { readFileSync, writeFileSync } from "fs";
import parser from "@asyncapi/parser";

await generate(
  "../../backend/todo-service/asyncapi.yaml",
  "../../backend/todo-service/app/todos/schemas"
);

await generate(
  "../../backend/notification-service/asyncapi.yaml",
  "../../backend/notification-service/NotificationService/Schemas"
);

async function generate(apiPath, outputPath) {
  console.log(`generating schemas from '${apiPath}'...`);

  const apiContent = readFileSync(apiPath, "utf8");

  const {
    _json: { channels },
  } = await parser.parse(apiContent);

  Object.entries(channels).forEach((e) => {
    const { publish } = e[1];

    if (!publish) return;

    const name = e[0];
    const data = JSON.stringify(publish?.message.payload, null, 2);

    const outputFilePath = `${outputPath}/${name}.json`;

    writeFileSync(outputFilePath, data);

    console.log(`    generated schema: ${outputFilePath}`);
  });
}
