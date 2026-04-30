"use client";

import Image from "next/image";
import { Bath, Building2, CheckCircle2, Mail, Phone, User } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";

export function PricingCard() {
  return (
    <Card className="border shadow-sm">
      <CardContent className="p-6">
        <div className="mb-6 flex items-center justify-between">
          <div className="flex items-baseline gap-1">
            <span className="text-3xl font-bold text-primary">$1,240</span>
            <span className="text-muted-foreground">/ month</span>
          </div>
          <span className="flex items-center gap-1 rounded-full bg-green-50 px-2 py-1 text-xs font-medium text-green-600">
            <CheckCircle2 className="h-3 w-3" />
            VERIFIED
          </span>
        </div>

        <div className="space-y-4 border-b pb-6">
          <div className="flex items-center justify-between">
            <div className="flex items-center gap-3 text-muted-foreground">
              <User className="h-4 w-4" />
              <span className="text-sm">Preferred Roommate</span>
            </div>
            <span className="text-sm font-medium">Any gender</span>
          </div>
          <div className="flex items-center justify-between">
            <div className="flex items-center gap-3 text-muted-foreground">
              <Bath className="h-4 w-4" />
              <span className="text-sm">Bathrooms</span>
            </div>
            <span className="text-sm font-medium">1 shared</span>
          </div>
          <div className="flex items-center justify-between">
            <div className="flex items-center gap-3 text-muted-foreground">
              <Building2 className="h-4 w-4" />
              <span className="text-sm">Availability</span>
            </div>
            <div className="flex items-center gap-2">
              <span className="text-sm font-medium">1 room open</span>
              <div className="flex gap-0.5">
                <div className="h-2 w-3 rounded-sm bg-primary" />
                <div className="h-2 w-3 rounded-sm bg-red-500" />
              </div>
            </div>
          </div>
        </div>

        <div className="border-b py-6">
          <p className="mb-4 text-xs text-muted-foreground">POSTED BY</p>
          <div className="flex items-center justify-between">
            <div className="flex items-center gap-3">
              <div className="relative h-10 w-10 overflow-hidden rounded-full bg-muted">
                <Image
                  src="https://api.dicebear.com/7.x/avataaars/svg?seed=SanctorHost"
                  alt="Poster avatar"
                  fill
                  className="object-cover"
                />
              </div>
              <div>
                <p className="text-sm font-medium">Maya K.</p>
                <p className="text-xs text-muted-foreground">
                  Community verified • 4.98 trust score
                </p>
              </div>
            </div>
            <div className="flex gap-2">
              <Button variant="outline" size="icon" className="h-8 w-8">
                <Mail className="h-4 w-4" />
              </Button>
              <Button variant="outline" size="icon" className="h-8 w-8">
                <Phone className="h-4 w-4" />
              </Button>
            </div>
          </div>
        </div>

        <div className="pt-6">
          <Button className="w-full bg-primary py-6 text-base text-primary-foreground hover:bg-primary/90">
            Message Poster
          </Button>
          <p className="mt-4 text-center text-xs text-muted-foreground">
            Questions stay connected to the community thread and listing history.
          </p>
        </div>
      </CardContent>
    </Card>
  );
}
