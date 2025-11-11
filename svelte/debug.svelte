<script>
  import Tooltip from "./_components/Tooltip.svelte";

  let tooltipComp = /** @type {Tooltip} */ (null);

  function onMouseEnterH3(/** @type {MouseEvent & { currentTarget: EventTarget & HTMLDivElement }} */ event) {

    const rect = event.currentTarget.getBoundingClientRect();

    tooltipComp.Show(rect);
  }

  function onMouseEnterH1(/** @type {MouseEvent & { currentTarget: EventTarget & HTMLDivElement }} */ event) {

    const rect = event.currentTarget.getBoundingClientRect();

    tooltipComp.Show(rect);
  }

  function onMouseEnterH2(/** @type {MouseEvent & { currentTarget: EventTarget & HTMLDivElement }} */ event) {

    const rect = event.currentTarget.getBoundingClientRect();

    tooltipComp.Show(rect);
  }
</script>

<main>
  <h1
  on:mouseenter={onMouseEnterH1}
  on:mouseleave={() => tooltipComp.Hide()}
  >Debug</h1>
  <div class="center">
    <h3
    class="h3-cool"
    on:mouseenter={onMouseEnterH3}
    on:mouseleave={() => tooltipComp.Hide()}
    >Hover me</h3>
    <Tooltip
      bind:this={tooltipComp}
    >
      <h3>Cihuyy HOVERED</h3>
    </Tooltip>
  </div>
  <h2
  on:mouseenter={onMouseEnterH2}
  on:mouseleave={() => tooltipComp.Hide()}
  >Bawah Banget</h2>
</main>

<style>
  .center {
    display: flex;
    height: 100%;
    width: 100%;
    justify-content: center;
    align-items: center;
  }

  .h3-cool {
    height: fit-content;
    width: fit-content;
  }
</style>