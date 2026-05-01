"use client";

import Link from "next/link";
import { Button } from "@/components/ui/button";

export function CommunityHeader() {
  return (
    <header className="sticky top-0 z-50 border-b border-border bg-card">
      <div className="mx-auto max-w-6xl px-4 sm:px-6 lg:px-8">
        <div className="flex h-16 items-center justify-between">
          <Link href="/communities" className="text-xl font-bold text-foreground">
            Sanctor
          </Link>

          <nav className="hidden items-center gap-8 md:flex">
            <Link href="/" className="text-sm text-muted-foreground transition-colors hover:text-foreground">
              Explore
            </Link>
            <Link href="/communities" className="border-b-2 border-foreground pb-0.5 text-sm font-medium text-foreground">
              Communities
            </Link>
            <Link href="/post-listings" className="text-sm text-muted-foreground transition-colors hover:text-foreground">
              Listings
            </Link>
          </nav>

          <div className="flex items-center gap-4">
            <Link href="#" className="text-sm text-muted-foreground transition-colors hover:text-foreground">
              Sign In
            </Link>
            <Button className="rounded-full bg-primary px-6 text-primary-foreground hover:bg-primary/90">
              Join Now
            </Button>
          </div>
        </div>
      </div>
    </header>
  );
}
