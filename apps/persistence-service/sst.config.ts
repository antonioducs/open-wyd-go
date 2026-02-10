import { SSTConfig } from "sst";
import { Function, Table } from "sst/constructs";

export default {
  config(_input) {
    return {
      name: "persistence-service",
      region: "us-east-1",
    };
  },
  stacks(app) {
    app.stack(function ({ stack }) {
      const table = new Table(stack, "WydTable", {
        fields: {
          pk: "string",
          sk: "string",
        },
        primaryIndex: { partitionKey: "pk", sortKey: "sk" },
      });

      const lambda = new Function(stack, "ApiHandler", {
        handler: "packages/functions/src/main.go",
        bind: [table],
        runtime: "go1.x",
        environment: {
          TABLE_NAME: table.tableName,
        }
      });

      stack.addOutputs({
        ApiEndpoint: lambda.url,
      });
    });
  },
} satisfies SSTConfig;
