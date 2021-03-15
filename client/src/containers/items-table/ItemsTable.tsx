import React, { useEffect, useState } from "react";
import { Notification, KIND } from "baseui/notification";
import { Table } from "baseui/table-semantic";
// import { Option } from "baseui/select";

import { Item } from "../../interfaces";
import { UseFetch } from "../../hooks/UseFetch";

export function ItemsTable(): JSX.Element {
  const { abortController, data, error, isLoading } = UseFetch<Item>(
    "/api/v1/items",
    undefined
  );
  const [rows, setRows] = useState<(string | number)[][]>([]);

  useEffect(() => {
    if (!isLoading && !error) {
      setRows(() =>
        data.map((item) => [
          item.id,
          item.itemName,
          item.quantity,
          item.category.categoryName,
        ])
      );
    }
  }, [isLoading, error, data]);

  useEffect(() => {
    return () => {
      abortController.abort();
    };
  }, [abortController]);

  return (
    <div>
      {error && (
        <Notification kind={KIND.negative}>{() => error.message}</Notification>
      )}
      <Table
        isLoading={isLoading}
        columns={["Id", "Name", "Quantity", "Category"]}
        data={rows}
      />
    </div>
  );
}
