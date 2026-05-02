import type { Metadata } from "next";
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
        {children}
        <FloatingMessageButton />
      </body>
    </html>
  );
}
