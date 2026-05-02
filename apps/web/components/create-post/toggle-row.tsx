"use client";

import { useState } from "react";

export function ToggleRow() {
  const [enabled, setEnabled] = useState(false);

  return (
    <div className="flex items-center justify-between gap-6 rounded-2xl border border-orange-100 bg-orange-50/40 px-6 py-5">
      <div>
        <p className="text-lg font-bold text-gray-900">Is this a sublet?</p>
        <p className="text-sm font-medium italic text-gray-500">
          Toggle if you are renting out your room while away.
        </p>
      </div>
      <button
        type="button"
        onClick={() => setEnabled((current) => !current)}
        className={`flex h-9 w-16 items-center rounded-full p-1 transition-colors ${
          enabled ? "bg-brand-orange" : "bg-gray-200"
        }`}
        aria-pressed={enabled}
      >
        <span
          className={`h-7 w-7 rounded-full bg-white shadow-sm transition-transform ${
            enabled ? "translate-x-7" : "translate-x-0"
          }`}
        />
      </button>
    </div>
  );
}
