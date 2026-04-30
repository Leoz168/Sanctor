"use client";

import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";

export function PremiumBanner() {
  return (
    <Card className="border-0 bg-primary text-primary-foreground">
      <CardContent className="p-6">
        <h3 className="mb-2 text-lg font-semibold">Community Context</h3>
        <p className="mb-4 text-sm opacity-90">
          This listing lives inside a housing community, so interested members can review prior posts, moderator notes, and shared move-in guidance before reaching out.
        </p>
        <Button variant="secondary" className="w-full bg-white text-primary hover:bg-white/90">
          View Community
        </Button>
      </CardContent>
    </Card>
  );
}
