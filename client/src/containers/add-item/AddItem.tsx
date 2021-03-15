import React, { useEffect, useState } from "react";
import { useStyletron } from "baseui";
import { Heading, HeadingLevel } from "baseui/heading";
import { Notification, KIND } from "baseui/notification";
import { FormControl } from "baseui/form-control";
import { Input } from "baseui/input";
import { Combobox } from "baseui/combobox";
import { Button } from "baseui/button";

import { UseFetch } from "../../hooks/UseFetch";
import { Category } from "../../interfaces";

export function AddItem() {
  const [css] = useStyletron();

  const {
    isLoading,
    error: categoriesError,
    data,
    abortController,
  } = UseFetch<Category>("/api/v1/categories");

  useEffect(() => {
    return () => {
      abortController.abort();
    };
  }, [abortController]);

  const [itemName, setItemName] = useState("");
  const [quantity, setQuantity] = useState<number>(1);
  const [categoryName, setCategoryName] = useState<string>("");
  const [error, setError] = useState<Error | null>(null);

  useEffect(() => {
    if (categoriesError) {
      setError(categoriesError);
    }
  }, [categoriesError]);

  const onSubmit = async (): Promise<void> => {
    try {
      const category = data.find(
        (category) => category.categoryName === categoryName
      );
      if (category) {
        await fetch(`/api/v1/categories/${category.id}/items`, {
          method: "POST",
          body: JSON.stringify({ itemName, quantity }),
        });
      }
    } catch (e) {
      setError(e as Error);
    }
  };

  return (
    <section className={css({ marginBottom: "16px", marginTop: "16px" })}>
      <HeadingLevel>
        <Heading styleLevel={5}>Add Item</Heading>
      </HeadingLevel>
      {error && (
        <Notification kind={KIND.negative}>{() => error.message}</Notification>
      )}
      <FormControl label={() => "Item Name"}>
        <Input
          size="compact"
          clearOnEscape={true}
          clearable={true}
          value={itemName}
          onChange={(e) => setItemName(e.currentTarget.value)}
          disabled={isLoading}
        />
      </FormControl>
      <FormControl label={() => "Quantity"}>
        <Input
          type="number"
          size="compact"
          min={1}
          max={1000}
          clearOnEscape={true}
          value={quantity}
          onChange={(e) => setQuantity(parseInt(e.currentTarget.value, 10))}
          disabled={isLoading}
        />
      </FormControl>
      <FormControl label={() => "Category"}>
        <Combobox
          size="compact"
          value={categoryName}
          options={data.map((category) => ({
            label: category.categoryName,
          }))}
          onChange={(e) => setCategoryName(e)}
          mapOptionToString={(option) => option.label}
        />
      </FormControl>

      <Button type="submit" onClick={onSubmit}>
        Add
      </Button>
    </section>
  );
}
