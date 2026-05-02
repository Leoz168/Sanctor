"use client";

import { useState } from "react";
import { Minus, Plus } from "lucide-react";

interface StepperProps {
  label: string;
  initialValue: number;
  min?: number;
  max?: number;
}

export function Stepper({ label, initialValue, min = 0, max = 10 }: StepperProps) {
  const [value, setValue] = useState(initialValue);

  return (
    <div>
      <p className="mb-2 text-center text-xs font-bold text-gray-400">{label}</p>
      <div className="flex items-center justify-center gap-5 rounded-2xl bg-brand-cream px-5 py-3">
        <button
          type="button"
          onClick={() => setValue((current) => Math.max(min, current - 1))}
          className="flex h-9 w-9 items-center justify-center rounded-lg bg-white text-brand-orange shadow-sm hover:bg-orange-50"
          aria-label={`Decrease ${label}`}
        >
          <Minus size={16} />
        </button>
        <span className="min-w-5 text-center text-xl font-bold text-gray-900">{value}</span>
        <button
          type="button"
          onClick={() => setValue((current) => Math.min(max, current + 1))}
          className="flex h-9 w-9 items-center justify-center rounded-lg bg-white text-brand-orange shadow-sm hover:bg-orange-50"
          aria-label={`Increase ${label}`}
        >
          <Plus size={16} />
        </button>
      </div>
    </div>
  );
}
