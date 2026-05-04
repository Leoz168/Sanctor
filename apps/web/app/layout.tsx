import type { Metadata } from "next";
import Script from "next/script";
import { FloatingMessageButton } from "@/components/messages/floating-message-button";
import "./globals.css";

export const metadata: Metadata = {
  title: "Rentling",
  description: "Community-first student housing platform",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>
        <Script src="https://accounts.google.com/gsi/client" strategy="afterInteractive" />
        {children}
        <FloatingMessageButton />
      </body>
    </html>
  );
}
