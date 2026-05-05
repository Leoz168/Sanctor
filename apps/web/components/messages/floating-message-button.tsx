"use client";

import { useRef, useState } from "react";
import type { PointerEvent as ReactPointerEvent } from "react";
import { Grip, Maximize2, MessageCircle, Minimize2, X } from "lucide-react";
import { Button } from "@/components/ui/button";
import { ContactsList } from "@/components/messages/contacts-list";
import { ChatWindow } from "@/components/messages/chat-window";

export type Contact = {
  id: string;
  name: string;
  avatar: string;
  online: boolean;
  lastMessage?: string;
  lastMessageTime?: string;
};

export type Message = {
  id: string;
  senderId: string;
  text: string;
  timestamp: string;
  type: "text" | "property";
  property?: {
    title: string;
    price: string;
    image: string;
  };
};

const mockContacts: Contact[] = [
  {
    id: "1",
    name: "Alex Rivera",
    avatar: "/images/community-1.jpg",
    online: true,
    lastMessage: "2:00 PM is perfect. Should I bring any...",
    lastMessageTime: "10:27 AM",
  },
  {
    id: "2",
    name: "Sarah Chen",
    avatar: "/images/community-4.jpg",
    online: true,
    lastMessage: "Thanks for the tour yesterday!",
    lastMessageTime: "Yesterday",
  },
  {
    id: "3",
    name: "Marcus Thompson",
    avatar: "/images/community-5.jpg",
    online: false,
    lastMessage: "I'll review the documents and get back...",
    lastMessageTime: "2 days ago",
  },
  {
    id: "4",
    name: "Elena Rodriguez",
    avatar: "/images/listing-4.jpg",
    online: false,
    lastMessage: "The apartment looks great!",
    lastMessageTime: "1 week ago",
  },
];

const mockMessages: Record<string, Message[]> = {
  "1": [
    {
      id: "m1",
      senderId: "1",
      text: "Hey! I just saw the listing for the downtown apartment. Is it still available for a viewing this Saturday?",
      timestamp: "10:24 AM",
      type: "text",
    },
    {
      id: "m2",
      senderId: "me",
      text: "Hi Alex! Yes, it is. We have a slot open at 2:00 PM. Does that work for you?",
      timestamp: "10:26 AM",
      type: "text",
    },
    {
      id: "m3",
      senderId: "1",
      text: "2:00 PM is perfect. Should I bring any specific documents with me?",
      timestamp: "10:27 AM",
      type: "text",
    },
    {
      id: "m4",
      senderId: "1",
      text: "",
      timestamp: "10:28 AM",
      type: "property",
      property: {
        title: "Skyline Loft - Unit 402",
        price: "$2,450 / mo",
        image: "/images/listing-1.jpg",
      },
    },
    {
      id: "m5",
      senderId: "me",
      text: "Just your ID for now. I'll send over the application forms after the viewing.",
      timestamp: "10:30 AM",
      type: "text",
    },
  ],
};

const defaultPanelSize = { width: 380, height: 600 };
const expandedPanelSize = { width: 760, height: 720 };
const minPanelSize = { width: 320, height: 420 };
const maxPanelSize = { width: 900, height: 760 };
const panelMargin = 24;
const dragSensitivity = 1;

type Point = {
  x: number;
  y: number;
};

type PanelSize = {
  width: number;
  height: number;
};

type PanelLayout = Point & PanelSize;
type ResizeDirection = "n" | "s" | "e" | "w" | "ne" | "nw" | "se" | "sw";

export function FloatingMessageButton() {
  const [isOpen, setIsOpen] = useState(false);
  const [isExpanded, setIsExpanded] = useState(false);
  const [selectedContact, setSelectedContact] = useState<Contact | null>(null);
  const [contacts] = useState<Contact[]>(mockContacts);
  const [layout, setLayout] = useState<PanelLayout>(() =>
    getDefaultLayout(defaultPanelSize),
  );

  const panelRef = useRef<HTMLDivElement | null>(null);
  const liveLayoutRef = useRef<PanelLayout>(layout);

  const applyLayout = (nextLayout: PanelLayout) => {
    liveLayoutRef.current = nextLayout;

    if (!panelRef.current) {
      return;
    }

    panelRef.current.style.left = `${nextLayout.x}px`;
    panelRef.current.style.top = `${nextLayout.y}px`;
    panelRef.current.style.width = `${nextLayout.width}px`;
    panelRef.current.style.height = `${nextLayout.height}px`;
  };

  const commitLayout = (nextLayout = liveLayoutRef.current) => {
    liveLayoutRef.current = nextLayout;
    setLayout(nextLayout);
    setIsExpanded(
      nextLayout.width > defaultPanelSize.width ||
        nextLayout.height > defaultPanelSize.height,
    );
  };

  const openMessages = () => {
    const nextLayout = getDefaultLayout(defaultPanelSize);

    liveLayoutRef.current = nextLayout;
    setLayout(nextLayout);
    setIsExpanded(false);
    setIsOpen(true);
  };

  const handleClose = () => {
    const nextLayout = getDefaultLayout(defaultPanelSize);

    setIsOpen(false);
    setIsExpanded(false);
    setSelectedContact(null);
    setLayout(nextLayout);
    liveLayoutRef.current = nextLayout;
  };

  const toggleExpanded = () => {
    const current = liveLayoutRef.current;
    const targetSize = isExpanded ? defaultPanelSize : expandedPanelSize;
    const target = normalizeSize(targetSize);
    const nextLayout = normalizeLayout({
      x: current.x + current.width - target.width,
      y: current.y + current.height - target.height,
      width: target.width,
      height: target.height,
    });

    applyLayout(nextLayout);
    commitLayout(nextLayout);
  };

  const startDrag = (event: ReactPointerEvent<HTMLButtonElement>) => {
    event.preventDefault();

    const startPointer = { x: event.clientX, y: event.clientY };
    const startLayout = liveLayoutRef.current;
    let nextLayout = startLayout;

    if (panelRef.current) {
      panelRef.current.style.willChange = "transform";
    }

    const movePanel = (moveEvent: PointerEvent) => {
      const dx = (moveEvent.clientX - startPointer.x) * dragSensitivity;
      const dy = (moveEvent.clientY - startPointer.y) * dragSensitivity;

      nextLayout = normalizeLayout({
        ...startLayout,
        x: startLayout.x + dx,
        y: startLayout.y + dy,
      });

      liveLayoutRef.current = nextLayout;

      if (!panelRef.current) {
        return;
      }

      panelRef.current.style.transform = `translate3d(${
        nextLayout.x - startLayout.x
      }px, ${nextLayout.y - startLayout.y}px, 0)`;
    };

    const stopDrag = () => {
      if (panelRef.current) {
        panelRef.current.style.transform = "";
        panelRef.current.style.willChange = "";
      }

      applyLayout(nextLayout);
      commitLayout(nextLayout);
      window.removeEventListener("pointermove", movePanel);
      window.removeEventListener("pointerup", stopDrag);
      window.removeEventListener("pointercancel", stopDrag);
    };

    window.addEventListener("pointermove", movePanel);
    window.addEventListener("pointerup", stopDrag);
    window.addEventListener("pointercancel", stopDrag);
  };

  const startResize = (
    direction: ResizeDirection,
    event: ReactPointerEvent<HTMLButtonElement>,
  ) => {
    event.preventDefault();
    event.stopPropagation();

    const startPointer = { x: event.clientX, y: event.clientY };
    const startLayout = liveLayoutRef.current;

    const resizePanel = (moveEvent: PointerEvent) => {
      const dx = moveEvent.clientX - startPointer.x;
      const dy = moveEvent.clientY - startPointer.y;

      applyLayout(resizeLayout(startLayout, direction, dx, dy));
    };

    const stopResize = () => {
      commitLayout();
      window.removeEventListener("pointermove", resizePanel);
      window.removeEventListener("pointerup", stopResize);
      window.removeEventListener("pointercancel", stopResize);
    };

    window.addEventListener("pointermove", resizePanel);
    window.addEventListener("pointerup", stopResize);
    window.addEventListener("pointercancel", stopResize);
  };

  return (
    <>
      {!isOpen && (
        <Button
          onClick={openMessages}
          className="fixed bottom-6 right-6 z-50 h-14 w-14 rounded-full bg-brand-orange text-white shadow-lg shadow-brand-orange/30 transition-transform hover:scale-105 hover:bg-orange-600"
          size="icon"
          aria-label="Open messages"
        >
          <MessageCircle className="h-6 w-6" />
        </Button>
      )}

      {isOpen && (
        <div
          ref={panelRef}
          className="fixed z-50 flex flex-col overflow-hidden rounded-2xl border border-border bg-card shadow-2xl"
          style={{
            left: layout.x,
            top: layout.y,
            width: layout.width,
            height: layout.height,
          }}
        >
          <div className="flex items-center justify-between border-b border-border bg-card px-4 py-3">
            <button
              type="button"
              onPointerDown={startDrag}
              className="flex min-w-0 flex-1 cursor-move touch-none select-none items-center gap-2 text-left"
              aria-label="Move messages"
            >
              <Grip className="h-4 w-4 shrink-0 text-muted-foreground" />
              <h2 className="font-semibold text-foreground">Messages</h2>
            </button>

            <div className="flex items-center gap-1">
              <Button
                variant="ghost"
                size="icon"
                onClick={toggleExpanded}
                className="h-8 w-8 rounded-full"
                aria-label={isExpanded ? "Shrink messages" : "Expand messages"}
              >
                {isExpanded ? (
                  <Minimize2 className="h-4 w-4" />
                ) : (
                  <Maximize2 className="h-4 w-4" />
                )}
              </Button>
              <Button
                variant="ghost"
                size="icon"
                onClick={handleClose}
                className="h-8 w-8 rounded-full"
                aria-label="Close messages"
              >
                <X className="h-4 w-4" />
              </Button>
            </div>
          </div>

          <div className="flex-1 overflow-hidden">
            {selectedContact ? (
              <ChatWindow
                contact={selectedContact}
                messages={mockMessages[selectedContact.id] || []}
                onBack={() => setSelectedContact(null)}
              />
            ) : (
              <ContactsList
                contacts={contacts}
                onSelectContact={setSelectedContact}
              />
            )}
          </div>

          <ResizeHandle direction="n" onResizeStart={startResize} />
          <ResizeHandle direction="s" onResizeStart={startResize} />
          <ResizeHandle direction="e" onResizeStart={startResize} />
          <ResizeHandle direction="w" onResizeStart={startResize} />
          <ResizeHandle direction="ne" onResizeStart={startResize} />
          <ResizeHandle direction="nw" onResizeStart={startResize} />
          <ResizeHandle direction="se" onResizeStart={startResize} />
          <ResizeHandle direction="sw" onResizeStart={startResize} />
        </div>
      )}
    </>
  );
}

interface ResizeHandleProps {
  direction: ResizeDirection;
  onResizeStart: (
    direction: ResizeDirection,
    event: ReactPointerEvent<HTMLButtonElement>,
  ) => void;
}

function ResizeHandle({ direction, onResizeStart }: ResizeHandleProps) {
  const className = getResizeHandleClassName(direction);

  return (
    <button
      type="button"
      className={className}
      onPointerDown={(event) => onResizeStart(direction, event)}
      aria-label={`Resize messages ${direction}`}
    />
  );
}

function getDefaultLayout(size: PanelSize): PanelLayout {
  if (typeof window === "undefined") {
    return { x: panelMargin, y: panelMargin, ...size };
  }

  const normalizedSize = normalizeSize(size);

  return normalizeLayout({
    x: window.innerWidth - normalizedSize.width - panelMargin,
    y: window.innerHeight - normalizedSize.height - panelMargin,
    ...normalizedSize,
  });
}

function normalizeSize(size: PanelSize): PanelSize {
  if (typeof window === "undefined") {
    return size;
  }

  return {
    width: Math.min(
      Math.max(minPanelSize.width, size.width),
      Math.min(maxPanelSize.width, window.innerWidth - panelMargin * 2),
    ),
    height: Math.min(
      Math.max(minPanelSize.height, size.height),
      Math.min(maxPanelSize.height, window.innerHeight - panelMargin * 2),
    ),
  };
}

function normalizeLayout(layout: PanelLayout): PanelLayout {
  const size = normalizeSize(layout);
  const maxX = Math.max(panelMargin, window.innerWidth - size.width - panelMargin);
  const maxY = Math.max(panelMargin, window.innerHeight - size.height - panelMargin);

  return {
    x: Math.min(Math.max(panelMargin, layout.x), maxX),
    y: Math.min(Math.max(panelMargin, layout.y), maxY),
    ...size,
  };
}

function resizeLayout(
  startLayout: PanelLayout,
  direction: ResizeDirection,
  dx: number,
  dy: number,
): PanelLayout {
  const nextLayout = { ...startLayout };

  if (direction.includes("e")) {
    nextLayout.width = startLayout.width + dx;
  }

  if (direction.includes("s")) {
    nextLayout.height = startLayout.height + dy;
  }

  if (direction.includes("w")) {
    nextLayout.width = startLayout.width - dx;
    nextLayout.x = startLayout.x + dx;
  }

  if (direction.includes("n")) {
    nextLayout.height = startLayout.height - dy;
    nextLayout.y = startLayout.y + dy;
  }

  const normalizedSize = normalizeSize(nextLayout);

  if (direction.includes("w")) {
    nextLayout.x = startLayout.x + startLayout.width - normalizedSize.width;
  }

  if (direction.includes("n")) {
    nextLayout.y = startLayout.y + startLayout.height - normalizedSize.height;
  }

  return normalizeLayout({
    ...nextLayout,
    ...normalizedSize,
  });
}

function getResizeHandleClassName(direction: ResizeDirection) {
  const base =
    "absolute z-10 touch-none rounded-full bg-transparent transition-colors hover:bg-brand-orange/20";

  const classes: Record<ResizeDirection, string> = {
    n: "left-5 right-5 top-0 h-2 cursor-n-resize",
    s: "bottom-0 left-5 right-5 h-2 cursor-s-resize",
    e: "bottom-5 right-0 top-5 w-2 cursor-e-resize",
    w: "bottom-5 left-0 top-5 w-2 cursor-w-resize",
    ne: "right-0 top-0 h-5 w-5 cursor-ne-resize",
    nw: "left-0 top-0 h-5 w-5 cursor-nw-resize",
    se: "bottom-0 right-0 h-6 w-6 cursor-se-resize",
    sw: "bottom-0 left-0 h-5 w-5 cursor-sw-resize",
  };

  return `${base} ${classes[direction]}`;
}
