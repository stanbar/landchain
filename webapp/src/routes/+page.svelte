<script lang="ts">
  import {  Button } from "carbon-components-svelte";
    
</script>


<div class="content">
    <img alt="logo" src="/logo.png" width={500}/>
    <h3>Wybierz co chcesz zrobic</h3>
    <br/><br/>
    <Button href="/create">Utwórz nowe zgłoszenie</Button>
    <br/>
    <Button href="/unapproved">Zobacz oczekujące zgłoszenia</Button>
    <br/>
    <Button href="/lands">Zobacz wszystkie dane księgi wieczystej</Button>

</div>
<style>
  .content {
    margin: 0 auto;
    display: flex;
    flex-direction: column;
    align-items: center;
  }
</style>