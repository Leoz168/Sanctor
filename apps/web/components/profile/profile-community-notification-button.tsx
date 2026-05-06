"use client";

import { useState } from "react";
import { Bell, BellOff } from "lucide-react";

export function ProfileCommunityNotificationButton() {
  const [isEnabled, setIsEnabled] = useState(true);
  const Icon = isEnabled ? Bell : BellOff;

  return (
    <button
      type="button"
      onClick={() => setIsEnabled((enabled) => !enabled)}
      className={`inline-flex h-10 w-10 items-center justify-center rounded-full border transition-all ${
        isEnabled
          ? "border-brand-orange bg-brand-orange text-white shadow-md shadow-brand-orange/20"
          : "border-orange-100 bg-white text-brand-orange hover:bg-brand-cream"
      }`}
      aria-label={
        isEnabled
          ? "Turn community notifications off"
          : "Turn community notifications on"
      }
      aria-pressed={isEnabled}
    >
      <Icon className="h-4 w-4" />
    </button>
  );
}
