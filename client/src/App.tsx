import React from "react";

import { Client as Styletron } from "styletron-engine-atomic";
import { Provider as StyletronProvider } from "styletron-react";
import { BaseProvider, LightTheme } from "baseui";
import { Heading, HeadingLevel } from "baseui/heading";

import { Container } from "./components/container";
import { AddItem } from "./containers/add-item";
import { ItemsTable } from "./containers/items-table";

const engine = new Styletron();

function App() {
  return (
    <StyletronProvider value={engine}>
      <BaseProvider theme={LightTheme}>
        <Container>
          <>
            <HeadingLevel>
              <Heading>Inventory</Heading>
            </HeadingLevel>
            <section>
              <AddItem />
              <ItemsTable />
            </section>
          </>
        </Container>
        {/* <Container>
          <AddItem categories={categories} items={items} />
        </Container> */}
      </BaseProvider>
    </StyletronProvider>
  );
}

export default App;
