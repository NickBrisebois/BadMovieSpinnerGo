const go = new Go();

WebAssembly.instantiateStreaming(fetch("static/spinner.wasm"), go.importObject).then((result) => {
  go.run(result.instance);
});

document.addEventListener("DOMContentLoaded", () => {
  const preventScrollOnCanvas = () => {
    const spinnerCanvas = document.querySelector("canvas");
    if (spinnerCanvas) {
      console.log("preventing ebitengine from taking scrollwheel events");
      spinnerCanvas.addEventListener(
        "wheel",
        (e) => {
          e.preventDefault();
          e.stopPropagation();
        },
        { capture: true, passing: false },
      );
    } else {
      setTimeout(preventScrollOnCanvas, 100);
    }
  };

  requestAnimationFrame(() => preventScrollOnCanvas());
});
