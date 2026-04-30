"use client";

import Link from "next/link";
import { Globe, Share2 } from "lucide-react";
import { Button } from "@/components/ui/button";

export function ListingFooter() {
  return (
    <footer className="border-t bg-card py-8">
      <div className="container mx-auto px-4">
        <div className="flex flex-col items-center justify-between gap-4 md:flex-row">
          <div>
            <p className="font-bold text-foreground">Sanctor</p>
            <p className="text-xs text-muted-foreground">
              © 2026 Sanctor. Community-first student housing discovery.
            </p>
          </div>

          <nav className="flex flex-wrap items-center justify-center gap-6">
            <Link href="#" className="text-xs text-muted-foreground underline hover:text-foreground">
              Terms of Service
            </Link>
            <Link href="#" className="text-xs text-muted-foreground underline hover:text-foreground">
              Privacy Policy
            </Link>
            <Link href="#" className="text-xs text-muted-foreground underline hover:text-foreground">
              Accessibility
            </Link>
            <Link href="#" className="text-xs text-muted-foreground underline hover:text-foreground">
              Contact Support
            </Link>
            <Link href="#" className="text-xs text-muted-foreground underline hover:text-foreground">
              Safety Guidelines
            </Link>
          </nav>

          <div className="flex items-center gap-2">
            <Button variant="ghost" size="icon" className="h-8 w-8">
              <Globe className="h-4 w-4" />
            </Button>
            <Button variant="ghost" size="icon" className="h-8 w-8">
              <Share2 className="h-4 w-4" />
            </Button>
          </div>
        </div>
      </div>
    </footer>
  );
}
