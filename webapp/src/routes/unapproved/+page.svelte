<script lang="ts">
    import { StructuredList, StructuredListHead, StructuredListRow, StructuredListCell, StructuredListBody, Link } from "carbon-components-svelte";
    import {State, type SignedRequest} from '$lib/types';
	import { grzechu, stachu, mikolaj, dominika } from "$lib/people";
  
    const firstLandID = "GD1G/00000123/1";
    const secondLandID = "WA1M/00005124/9";

    const requests: SignedRequest[] = [
        {
        id: "",
        signatures: ["sig1", "sig2"],
        state: State.WAITING_FOR_APPROVAL,
        landID: firstLandID,
        creationTime: "2022-10-16 03:13:11",
        parties: [grzechu, stachu],
        intermediary: {
             name: "Jakub Kowalski",
             pesel: "49012345678",
             publicKey: "",
         },
        newOwner: stachu,
        previousNotarialAct: {
            landID: firstLandID,
            notary: mikolaj,
            creationDate: "",
            description: "Transfer własności działki.",
            owner: grzechu
        },
        newNotarialAct: {
            landID: firstLandID,
            notary: mikolaj,
            creationDate: "",
            description: "Odziedziczenie działki w spadku.",
            owner: stachu,
        },
    },
    {
        id: "",
        signatures: ["sig1", "sig2"],
        state: State.WAITING_FOR_APPROVAL,
        landID: secondLandID,
        creationTime: "2022-10-16 06:13:11",
        parties: [dominika, mikolaj],
        intermediary: grzechu,
        newOwner: mikolaj,
        previousNotarialAct: {
            landID: secondLandID,
            notary: stachu,
            creationDate: "",
            description: "Odziedziczenie działki w spadku.",
            owner: dominika
        },
        newNotarialAct: {
            landID: secondLandID,
            notary: stachu,
            creationDate: "",
            description: "Transfer własności działki.",
            owner: mikolaj,
        },
    }
]; 
    
    const onConfirm = (request: SignedRequest) => {
        alert(`confirming ${request.id}`);
    }
    const onReject = (request: SignedRequest) => {
        alert(`rejecting ${request.id}`);
    }
    
</script>


<div class="content">
    <img alt="logo" src="/logo.png" width={500}/>
    <StructuredList>
        <StructuredListHead>
            <StructuredListRow head>
                <StructuredListCell head>Data utworzenia</StructuredListCell>
                <StructuredListCell head>Numer księgi wieczystej</StructuredListCell>
                <StructuredListCell head>Sprzedawca</StructuredListCell>
                <StructuredListCell head>Kupujący</StructuredListCell>
                <StructuredListCell head>Pośrednik</StructuredListCell>
                <StructuredListCell head>Poprzedni akt notarialny</StructuredListCell>
                <StructuredListCell head>Nowy akt notarialny</StructuredListCell>
                <StructuredListCell head>Status</StructuredListCell>
                <StructuredListCell head>Akcja</StructuredListCell>
            </StructuredListRow>
        </StructuredListHead>
        <StructuredListBody>
            {#each requests as request}
                <StructuredListRow>
                    <StructuredListCell>
                        {request.creationTime}
                    </StructuredListCell>
                    <StructuredListCell>
                        {request.landID}
                    </StructuredListCell>
                    <StructuredListCell>
                        {request.parties[0].name} - {request.parties[0].pesel}
                    </StructuredListCell>
                    <StructuredListCell>
                        {request.parties[1].name} - {request.parties[1].pesel}
                    </StructuredListCell>
                    <StructuredListCell>
                        {request.intermediary.name} - {request.intermediary.pesel}
                    </StructuredListCell>
                    <StructuredListCell>
                        {request.previousNotarialAct.description}
                    </StructuredListCell>
                    <StructuredListCell>
                        {request.newNotarialAct.description}
                    </StructuredListCell>
                    <StructuredListCell>
                        {request.state}
                    </StructuredListCell>
                    <StructuredListCell>
                        <Link on:click={() => onConfirm(request)}>Zatwierdź</Link>
                        <Link on:click={() => onConfirm(request)}>Zapłać</Link>
                        <Link on:click={() => onReject(request)} style="color: red;">Odrzuć</Link>
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
    padding: 100px;
    flex-direction: column;
    align-items: center;
  }
</style>