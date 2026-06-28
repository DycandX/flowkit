#!/bin/bash
# generate_workflow.sh — Generate WORKFLOW.md from template
# Usage:
#   Interactive: bash generate_workflow.sh
#   With args:   bash generate_workflow.sh "Project Name" "main" "en" "npm install" "npm run dev" "npm run build" "npm run lint"
#
# Arguments (positional):
#   1 = PROJECT_NAME
#   2 = MAIN_BRANCH       (default: main)
#   3 = LANGUAGE          (id or en, default: id)
#   4 = INSTALL_COMMAND   (default: npm install)
#   5 = DEV_COMMAND       (default: npm run dev)
#   6 = BUILD_COMMAND     (default: npm run build)
#   7 = LINT_COMMAND      (default: npm run lint)

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# ─── Interactive mode ──────────────────────────────────────────────────
if [ $# -lt 1 ]; then
  echo "╔══════════════════════════════════════════════╗"
  echo "║   Generate WORKFLOW.md from Template         ║"
  echo "╚══════════════════════════════════════════════╝"
  echo ""
  
  read -r -p "Project name       : " PROJECT_NAME
  read -r -p "Main branch        [main]: " MAIN_BRANCH_INPUT
  read -r -p "Language (id/en)   [id]: " LANG_INPUT
  read -r -p "Install command    [npm install]: " INSTALL_INPUT
  read -r -p "Dev command        [npm run dev]: " DEV_INPUT
  read -r -p "Build command      [npm run build]: " BUILD_INPUT
  read -r -p "Lint command       [npm run lint]: " LINT_INPUT
  
  MAIN_BRANCH="${MAIN_BRANCH_INPUT:-main}"
  LANG="${LANG_INPUT:-id}"
  INSTALL_CMD="${INSTALL_INPUT:-npm install}"
  DEV_CMD="${DEV_INPUT:-npm run dev}"
  BUILD_CMD="${BUILD_INPUT:-npm run build}"
  LINT_CMD="${LINT_INPUT:-npm run lint}"
else
  PROJECT_NAME="$1"
  MAIN_BRANCH="${2:-main}"
  LANG="${3:-id}"
  INSTALL_CMD="${4:-npm install}"
  DEV_CMD="${5:-npm run dev}"
  BUILD_CMD="${6:-npm run build}"
  LINT_CMD="${7:-npm run lint}"
fi

# ─── Validate ─────────────────────────────────────────────────────────
if [ -z "$PROJECT_NAME" ]; then
  echo "❌ Error: PROJECT_NAME is required."
  exit 1
fi

if [ "$LANG" != "id" ] && [ "$LANG" != "en" ]; then
  echo "❌ Error: Language must be 'id' or 'en'."
  exit 1
fi

TEMPLATE_FILE="$SCRIPT_DIR/template_workflow.$LANG.md"

if [ ! -f "$TEMPLATE_FILE" ]; then
  echo "❌ Error: Template file not found: $TEMPLATE_FILE"
  echo "   Expected at: $TEMPLATE_FILE"
  exit 1
fi

# ─── Get stack info ───────────────────────────────────────────────────
echo ""
echo "Detecting project stack..."
STACK=""

if [ -f "package.json" ]; then
  if grep -qi '"next"' package.json 2>/dev/null; then
    STACK="Next.js"
  elif grep -qi '"react"' package.json 2>/dev/null; then
    STACK="React"
  elif grep -qi '"nuxt"' package.json 2>/dev/null; then
    STACK="Nuxt.js"
  elif grep -qi '"vue"' package.json 2>/dev/null; then
    STACK="Vue"
  else
    STACK="Node.js"
  fi
elif [ -f "composer.json" ]; then
  STACK="Laravel / PHP"
elif [ -f "requirements.txt" ] || [ -f "pyproject.toml" ]; then
  STACK="Python"
elif [ -f "Cargo.toml" ]; then
  STACK="Rust"
elif [ -f "go.mod" ]; then
  STACK="Go"
else
  STACK="Unknown"
fi

echo "   → Detected: $STACK"

# ─── Get root path ────────────────────────────────────────────────────
CURRENT_DIR="$(pwd)"
PROJECT_DIRNAME="$(basename "$CURRENT_DIR")"

echo ""
echo "Generating WORKFLOW.md..."
echo "   Project  : $PROJECT_NAME"
echo "   Branch   : $MAIN_BRANCH"
echo "   Language : $LANG"
echo "   Stack    : $STACK"
echo "   Path     : $CURRENT_DIR"

# ─── Generate file ────────────────────────────────────────────────────
OUTPUT_FILE="$CURRENT_DIR/WORKFLOW.md"

sed -e "s/{{PROJECT_NAME}}/$PROJECT_NAME/g" \
    -e "s/{{MAIN_BRANCH}}/$MAIN_BRANCH/g" \
    -e "s|{{ROOT_PATH}}|$CURRENT_DIR|g" \
    -e "s|{{STACK}}|$STACK|g" \
    -e "s|{{INSTALL_COMMAND}}|$INSTALL_CMD|g" \
    -e "s|{{DEV_COMMAND}}|$DEV_CMD|g" \
    -e "s|{{BUILD_COMMAND}}|$BUILD_CMD|g" \
    -e "s|{{LINT_COMMAND}}|$LINT_CMD|g" \
    "$TEMPLATE_FILE" > "$OUTPUT_FILE"

if [ $? -eq 0 ]; then
  echo ""
  echo "✅ Done! File created:"
  echo "   $OUTPUT_FILE"
  echo ""
  echo "Next step:"
  echo "   1. Add to your AGENTS.md:"
  echo '      > Read `WORKFLOW.md` for branching strategy, commit convention, and workflow.'
  echo "   2. Create develop branch:"
  echo "      git checkout -b develop"
  echo "      git push origin develop"
else
  echo "❌ Error: Failed to generate WORKFLOW.md"
  exit 1
fi
