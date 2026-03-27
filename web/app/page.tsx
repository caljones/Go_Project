"use client";

import { useEffect, useRef, useState } from "react";

declare global {
  interface Window {
    Go: new () => {
      importObject: WebAssembly.Imports;
      run: (instance: WebAssembly.Instance) => Promise<void>;
    };
    __goOutput: (text: string) => void;
    __goDone: () => void;
    __toursUrl: string;
  }
}

export default function Home() {
  const terminalRef = useRef<HTMLDivElement>(null);
  const [isDone, setIsDone] = useState(false);

  useEffect(() => {
    const basePath = process.env.NEXT_PUBLIC_BASE_PATH || '';

    const initWasm = async () => {
      if (!terminalRef.current) return;

      const { Terminal } = await import("@xterm/xterm");
      const { FitAddon } = await import("@xterm/addon-fit");
      await import("@xterm/xterm/css/xterm.css");

      const terminal = new Terminal({
        cursorBlink: true,
        theme: {
          background: "#000000",
          foreground: "#39ff14",
          cursor: "#39ff14",
        },
        fontFamily: "'Courier New', Courier, monospace",
        fontSize: 14,
        convertEol: false,
      });

      const fitAddon = new FitAddon();
      terminal.loadAddon(fitAddon);
      terminal.open(terminalRef.current);
      fitAddon.fit();

      window.addEventListener("resize", () => fitAddon.fit());

      // Expose globals for the WASM binary
      window.__toursUrl = `${window.location.origin}${basePath}/go/tours.json`;
      window.__goOutput = (text: string) => {
        // Convert bare \n to \r\n for proper xterm line breaks
        terminal.write(text.replace(/\n/g, "\r\n"));
      };

      window.__goDone = () => {
        setIsDone(true);
      };

      // Load and run the WASM — go.run() never resolves (select{} in main_wasm.go)
      const go = new window.Go();
      const result = await WebAssembly.instantiateStreaming(
        fetch(`${basePath}/go/main.wasm`),
        go.importObject
      );
      go.run(result.instance);
    };

    const script = document.createElement("script");
    script.src = `${basePath}/wasm_exec.js`;
    script.onload = initWasm;
    document.body.appendChild(script);

    return () => { document.body.removeChild(script); };
  }, []);

  return (
    <>
      <div id="terminal-container" ref={terminalRef} />
      <div id="links" style={{ display: isDone ? "flex" : "none" }}>
        <a href="mailto:cnjones7@ncsu.edu">cnjones7@ncsu.edu</a>
        <a
          href="https://linkedin.com/in/calvin-n-jones"
          target="_blank"
          rel="noopener noreferrer"
        >
          linkedin.com/in/calvin-n-jones
        </a>
      </div>
    </>
  );
}
