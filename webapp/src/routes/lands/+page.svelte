<script lang="ts">
	import { grzechu, mikolaj, stachu, } from "$lib/people";
	import type { Land } from "$lib/types";
    import { StructuredList, StructuredListHead, StructuredListRow, StructuredListCell, StructuredListBody, Link } from "carbon-components-svelte";
  
    const lands: Land[] = [
        {
        id: "",
        address: "ul. Kolejowa 1, 00-000 Warszawa",
        owner: grzechu,
        area: 123,
        mortgage: {
            creditor: "BRE BANK S.A",
            amount: 500_000,
        },
        notarialAct: {
            landID: "GD1G/00000123/1",
            notary: mikolaj,
            creationDate: "2001-10-16 04:13:11",
            description: "Transfer własności działki.",
            owner: grzechu
        },
    },
    {
        id: "",
        address: "ul. Kolejowa 2, 00-000 Warszawa",
        owner: stachu,
        area: 321,
        mortgage: {
            creditor: "PKO BP S.A",
            amount: 2_100_000,
        },
        notarialAct: {
            landID: "WA1M/00005124/9",
            notary: mikolaj,
            creationDate: "2009-03-23 06:15:11",
            description: "Odziedziczenie w ramach spadku.",
            owner: stachu
        },
    }
    ]; 
</script>


<div class="content">
    <img alt="logo" src="/logo.png" width={500}/>
    <StructuredList>
        <StructuredListHead>
            <StructuredListRow head>
                <StructuredListCell head>Adres</StructuredListCell>
                <StructuredListCell head>Data utworzenia</StructuredListCell>
                <StructuredListCell head>Numer księgi wieczystej</StructuredListCell>
                <StructuredListCell head>Nazwa Własciciela</StructuredListCell>
                <StructuredListCell head>PESEL Własciciela</StructuredListCell>
                <StructuredListCell head>Powierzchnia</StructuredListCell>
                <StructuredListCell head>Kredytodawca hipoteki</StructuredListCell>
                <StructuredListCell head>Wartośc hipoteki</StructuredListCell>
            </StructuredListRow>
        </StructuredListHead>
        <StructuredListBody>
            {#each lands as land}
                <StructuredListRow>
                    <StructuredListCell>
                        {land.address}
                    </StructuredListCell>
                    <StructuredListCell>
                        {land.notarialAct.creationDate}
                    </StructuredListCell>
                    <StructuredListCell>
                        {land.notarialAct.landID}
                    </StructuredListCell>
                    <StructuredListCell>
                        {land.owner.name}
                    </StructuredListCell>
                    <StructuredListCell>
                        {land.owner.pesel}
                    </StructuredListCell>
                    <StructuredListCell>
                        {land.area}m&#178;
                    </StructuredListCell>
                    <StructuredListCell>
                        {land.mortgage.creditor}
                    </StructuredListCell>
                    <StructuredListCell>
                        {land.mortgage.amount}zł
                    </StructuredListCell>
                </StructuredListRow>
            {/each}
        </StructuredListBody>
    </StructuredList>

</div>
<style>
  .content {
    margin: 0 auto;
    display: flex;
    flex-direction: column;
    align-items: center;
  }
</style>