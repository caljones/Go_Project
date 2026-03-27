import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Calvin Jones — Go Demo",
  description: "A Go CLI program running in the browser via WebAssembly",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
