"use client";

import { Car, Dumbbell, Wifi, WashingMachine } from "lucide-react";

const amenities = [
  { icon: Wifi, label: "High-speed WiFi" },
  { icon: Car, label: "Street Parking Nearby" },
  { icon: Dumbbell, label: "Gym in Building" },
  { icon: WashingMachine, label: "In-unit Laundry" },
];

export function Amenities() {
  return (
    <section>
      <h2 className="mb-6 text-xl font-semibold">Amenities</h2>
      <div className="flex flex-wrap gap-4">
        {amenities.map((amenity) => (
          <div
            key={amenity.label}
            className="flex items-center gap-3 rounded-lg bg-secondary/50 px-4 py-3"
          >
            <div className="flex h-8 w-8 items-center justify-center rounded-lg bg-primary/10">
              <amenity.icon className="h-4 w-4 text-primary" />
            </div>
            <span className="text-sm font-medium">{amenity.label}</span>
          </div>
        ))}
      </div>
    </section>
  );
}
