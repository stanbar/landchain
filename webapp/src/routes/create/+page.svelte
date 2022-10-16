<script lang="ts">
  import { TextInput, TextArea, Grid, Row, Column, Button } from "carbon-components-svelte";
  import Add from "carbon-icons-svelte/lib/Add.svelte";

  export let request = {
      landID: '',
      buyer: {
          name: '',
          pesel: '',
      },
      seller: {
          name: '',
          pesel: '',
      },
      intermediary: {
          name: '',
          pesel: '',
      },
      notarialAct: {
        notary: {
          name: '',
          pesel: '',
        },
        description: '',
      }
  };
  
  const onSubmit = () => {
    console.log({request});

  };
</script>
<div class="content">
  <img alt="logo" src="/logo.png" width={500}/>
  <h3>Podstawowe dane</h3>
  <TextInput labelText="Numer wpisu w księdze wieczystej" bind:value={request.landID} />
  <br/>
  
  <Grid>
    <Row>
      <Column>
        <h3>Kupujący</h3>
        <TextInput labelText="Nazwa" bind:value={request.buyer.name} />
        <br/>
        <TextInput labelText="Numer PESEL" bind:value={request.buyer.pesel} />
        <br/><br/>
      </Column>
      <Column>
        <h3>Sprzedający</h3>
        <TextInput labelText="Nazwa" bind:value={request.seller.name} />
        <br/>
        <TextInput labelText="Numer PESEL" bind:value={request.seller.pesel} />
        <br/><br/>
      </Column>
    </Row>
  </Grid>

  
  <h3>Pośrednik</h3>
  <TextInput labelText="Nazwa" bind:value={request.intermediary.name} />
  <br/>
  <TextInput labelText="Numer PESEL" bind:value={request.intermediary.pesel} />
  <br/><br/>

  <h3>Akt Notarialny</h3>
  <br/>
  <TextInput labelText="Nazwa notariusza" bind:value={request.notarialAct.notary.name} />
  <br/>
  <TextInput labelText="Numer PESEL notariusza" bind:value={request.notarialAct.notary.pesel} />
  <br/>
  <TextArea labelText="Opis" bind:value={request.notarialAct.description}/>
  <br/><br/>

  <Button icon={Add} on:click={onSubmit}>Dodaj</Button>
</div>
<style>
  .content {
    margin: 0 auto;
    display: flex;
    flex-direction: column;
    align-items: center;
  }
</style>
